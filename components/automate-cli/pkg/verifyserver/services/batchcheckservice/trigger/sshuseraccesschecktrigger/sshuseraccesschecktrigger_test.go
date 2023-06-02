package sshuseraccesschecktrigger

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/lib/logger"
	"github.com/stretchr/testify/assert"
)

// mockTransport is a mock implementation of the http.RoundTripper interface
type mockTransport struct{}

// RoundTrip returns an error for every request
func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("mock error")
}

const (
	sshCheckSuccessResp = `{
		"status": "SUCCESS",
		"result": {
		  "passed": true,
		  "checks": [
			{
			  "title": "SSH user accessible",
			  "passed": true,
			  "success_msg": "SSH user is accessible for the node: <node_ip>",
			  "error_msg": "",
			  "resolution_msg": ""
			},
			{
			  "title": "Sudo password valid",
			  "passed": true,
			  "success_msg": "SSH user sudo password is valid for the node: <node_ip>",
			  "error_msg": "",
			  "resolution_msg": ""
			}
		  ]
		}
	  }`
	sshCheckFailureResp = `{
		"status": "SUCCESS",
		"result": {
		  "passed": false,
		  "checks": [
			{
			  "title": "SSH user unaccessible",
			  "passed": false,
			  "success_msg": "",
			  "error_msg": "SSH user is unaccessible for the node with IP <node_ip>",
			  "resolution_msg": "Give SSH access to the user with the give key on the node: <node_ip>"
			},
			{
			  "title": "Sudo password invalid",
			  "passed": false,
			  "success_msg": "",
			  "error_msg": "SSH user sudo password is invalid for the node with IP <node_ip>",
			  "resolution_msg": "Ensure you have provided the correct sudo password and the user has sudo access on the node: <node_ip>"
			}
		  ]
		}
	  }`
)

func GetRequestJson() models.Config {
	ipConfig := models.Config{}

	json.Unmarshal([]byte(`{
		  "ssh_user": {
			"user_name": "ubuntu",
			"private_key": "test_key",
			"sudo_password": "test@123"
		  },
		  "arch": "existing_nodes",
		  "backup": {
			"file_system": {
			  "mount_location": "/mnt/automate_backups"
			}
		  },
		  "hardware": {
			"automate_node_count": 1,
			"automate_node_ips": [
			  "1.2.3.4"
			],
			"chef_infra_server_node_count": 1,
			"chef_infra_server_node_ips": [
			  "5.6.7.8"
			],
			"postgresql_node_count": 1,
			"postgresql_node_ips": [
			  "9.10.11.12"
			],
			"opensearch_node_count": 1,
			"opensearch_node_ips": [
			  "14.15.16.17"
			]
		  },
		  "certificate": {
			"fqdn":"my_fqdn",
			"root_cert":"---- VALID ROOT CA ----",
			"nodes":[
				{
					"ip":"1.2.3.4",
					"cert":"---- VALID NODE CERT ----",
					"key":"---- VALID PRIVATE KEY ----",
					"admin_key":"---- VALID ADMIN PRIVATE KEY ----",
					"admin_cert":"---- VALID ADMIN CERT ----"
				},
				{
					"ip":"5.6.7.8",
					"cert":"---- VALID NODE CERT ----",
					"key":"---- VALID PRIVATE KEY ----",
					"admin_key":"---- VALID ADMIN PRIVATE KEY ----",
					"admin_cert":"---- VALID ADMIN CERT ----"
				},
				{
					"ip":"9.10.11.12",
					"cert":"---- VALID NODE CERT ----",
					"key":"---- VALID PRIVATE KEY ----",
					"admin_key":"---- VALID ADMIN PRIVATE KEY ----",
					"admin_cert":"---- VALID ADMIN CERT ----"
				},
				{
					"ip":"14.15.16.17",
					"cert":"---- VALID NODE CERT ----",
					"key":"---- VALID PRIVATE KEY ----",
					"admin_key":"---- VALID ADMIN PRIVATE KEY ----",
					"admin_cert":"---- VALID ADMIN CERT ----"
				}
			]
		  }
		}`), &ipConfig)
	return ipConfig
}

func startMockServerOnCustomPort(mockServer *httptest.Server, port string) error {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		return err
	}
	mockServer.Listener = l
	mockServer.Start()
	return nil
}

func TestNewSshUserAccessCheck(t *testing.T) {
	testPort := "1234"
	cc := NewSshUserAccessCheck(logger.NewTestLogger(), testPort)
	assert.NotNil(t, cc)
	assert.NotNil(t, cc.log)
	assert.Equal(t, constants.LOCAL_HOST_URL, cc.host)
	assert.Equal(t, testPort, cc.port)
}

func TestSshUserAccessCheck_Run(t *testing.T) {
	t.Run("Returns OK", func(t *testing.T) {
		request := GetRequestJson()
		//starting the mock server on custom port
		mockServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r.Body
			//not handling error because of test file
			reader, _ := io.ReadAll(req)
			var actualRequest models.SShUserRequest
			json.Unmarshal([]byte(reader), &actualRequest)

			sshRequest := GetSshRequest()
			assert.NotNil(t, actualRequest)
			assert.Equal(t, actualRequest.PrivateKey, sshRequest.PrivateKey)
			assert.Equal(t, actualRequest.Username, sshRequest.Username)
			assert.Equal(t, actualRequest.SudoPassword, sshRequest.SudoPassword)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(sshCheckSuccessResp))
		}))
		err := startMockServerOnCustomPort(mockServer, "1124")
		assert.NoError(t, err)
		defer mockServer.Close()

		cc := NewSshUserAccessCheck(logger.NewTestLogger(), "1124")

		finalResp := cc.Run(request)
		totalIps := request.Hardware.AutomateNodeCount + request.Hardware.ChefInfraServerNodeCount + request.Hardware.PostgresqlNodeCount + request.Hardware.OpenSearchNodeCount
		assert.Equal(t, totalIps, len(finalResp))

		for _, resp := range finalResp {
			assert.NotEmpty(t, resp.Result.Checks)
			if resp.Host == "14.15.16.17" {
				triggerResp := resp
				assert.Equal(t, "SUCCESS", triggerResp.Status)
				assert.NotEmpty(t, triggerResp.Result)
				assert.Equal(t, resp.Result.Passed, true)
				assert.Equal(t, 2, len(triggerResp.Result.Checks))
				assert.Equal(t, "SSH user accessible", triggerResp.Result.Checks[0].Title)
				assert.Equal(t, true, triggerResp.Result.Checks[0].Passed)
				assert.Equal(t, "SSH user is accessible for the node: <node_ip>", triggerResp.Result.Checks[0].SuccessMsg)
				assert.Nil(t, triggerResp.Result.Error)

			} else {
				assert.Equal(t, resp.Result.Passed, true)
			}
		}
	})

	t.Run("Failure response", func(t *testing.T) {
		//starting the mock server on custom port
		mockServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r.Body
			//not handling error because of test file
			reader, _ := io.ReadAll(req)
			var actualRequest models.SShUserRequest
			json.Unmarshal([]byte(reader), &actualRequest)

			sshRequest := GetSshRequest()
			assert.NotNil(t, actualRequest)
			assert.Equal(t, actualRequest.PrivateKey, sshRequest.PrivateKey)
			assert.Equal(t, actualRequest.Username, sshRequest.Username)
			assert.Equal(t, actualRequest.SudoPassword, sshRequest.SudoPassword)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(sshCheckFailureResp))
		}))
		err := startMockServerOnCustomPort(mockServer, "1235")
		assert.NoError(t, err)
		defer mockServer.Close()

		cc := NewSshUserAccessCheck(logger.NewTestLogger(), "1235")
		request := GetRequestJson()
		finalResp := cc.Run(request)
		totalIps := request.Hardware.AutomateNodeCount + request.Hardware.ChefInfraServerNodeCount + request.Hardware.PostgresqlNodeCount + request.Hardware.OpenSearchNodeCount
		assert.Equal(t, totalIps, len(finalResp))

		for _, resp := range finalResp {
			assert.NotEmpty(t, resp.Result.Checks)
			if resp.Host == "14.15.16.17" {
				triggerResp := resp
				assert.Equal(t, "SUCCESS", triggerResp.Status)
				assert.NotEmpty(t, triggerResp.Result)
				assert.Equal(t, resp.Result.Passed, false)
				assert.Equal(t, 2, len(triggerResp.Result.Checks))
				assert.Equal(t, "Sudo password invalid", triggerResp.Result.Checks[1].Title)
				assert.Equal(t, false, triggerResp.Result.Checks[1].Passed)
				assert.Equal(t, "SSH user sudo password is invalid for the node with IP <node_ip>", triggerResp.Result.Checks[1].ErrorMsg)
				assert.Equal(t, "Ensure you have provided the correct sudo password and the user has sudo access on the node: <node_ip>", triggerResp.Result.Checks[1].ResolutionMsg)
				assert.Empty(t, triggerResp.Result.Checks[1].SuccessMsg)
				assert.Nil(t, triggerResp.Result.Error)

			} else {
				assert.Equal(t, resp.Result.Passed, false)
			}
		}
	})

	t.Run("Returns error", func(t *testing.T) {
		mockServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`invalid JSON`))
		}))
		err := startMockServerOnCustomPort(mockServer, "1236")
		assert.NoError(t, err)
		defer mockServer.Close()

		cc := NewSshUserAccessCheck(logger.NewTestLogger(), "1236")
		request := GetRequestJson()
		finalResp := cc.Run(request)
		totalIps := request.Hardware.AutomateNodeCount + request.Hardware.ChefInfraServerNodeCount + request.Hardware.PostgresqlNodeCount + request.Hardware.OpenSearchNodeCount
		assert.Equal(t, totalIps, len(finalResp))

		for _, resp := range finalResp {
			assert.NotNil(t, resp.Result.Error)
			assert.Empty(t, resp.Result.Checks)
			assert.Equal(t, resp.Result.Passed, false)
		}
	})

}

func TestGetSShUserAPIRquest(t *testing.T) {
	ip := "1.2.3.4"

	expected := GetSshRequest()

	actual := getSShUserAPIRquest(ip, GetRequestJson().SSHUser)

	assert.Equal(t, expected, actual)

}

func GetSshRequest() models.SShUserRequest {
	request := models.SShUserRequest{
		IP:           "1.2.3.4",
		Username:     "ubuntu",
		PrivateKey:   "test_key",
		SudoPassword: "test@123",
	}
	return request
}