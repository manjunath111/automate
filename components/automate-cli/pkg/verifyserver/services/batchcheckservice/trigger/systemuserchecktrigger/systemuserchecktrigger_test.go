package systemuserchecktrigger

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/lib/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	systemUser = `{
		"status": "SUCCESS",
		"result": {
			"passed": true,
			"checks": [
				{
					"title": "User creation/validation check",
					"passed": true,
					"success_msg": "User is created or found successfully",
					"error_msg": "",
					"resolution_msg": ""
				},
				{
					"title": "User and group mapping successfully",
					"passed": true,
					"success_msg": "User and group mapping successful",
					"error_msg": "",
					"resolution_msg": ""
				}
			]
		}
	}`
)

func TestSystemUserCheck_Run(t *testing.T) {
	t.Run("System User Check", func(t *testing.T) {
		// Create a dummy server
		server, host, port := createDummyServer(t, http.StatusOK)
		defer server.Close()

		// Test data
		config := &models.Config{
			Hardware: &models.Hardware{
				AutomateNodeCount: 1,
				AutomateNodeIps:   []string{host},
			},
		}

		suc := NewSystemUserCheck(logger.NewLogrusStandardLogger(), port)
		ctr := suc.Run(config)

		// Assert the expected result
		require.Len(t, ctr, 1)
		require.Nil(t, ctr[0].Result.Error)
		require.Len(t, ctr[0].Result.Checks, 2)
		require.Equal(t, "", ctr[0].Result.Check)

		checkResp := ctr[0].Result.Checks[1]

		assert.Equal(t, "User and group mapping successfully", checkResp.Title)
		assert.Equal(t, true, checkResp.Passed)
		assert.Equal(t, "User and group mapping successful", checkResp.SuccessMsg)
		assert.Equal(t, "", checkResp.ErrorMsg)
		assert.Equal(t, "", checkResp.ResolutionMsg)
	})

	t.Run("Failed User Check", func(t *testing.T) {
		// Create a dummy server
		server, host, port := createDummyServer(t, http.StatusInternalServerError)
		defer server.Close()

		// Test data
		config := &models.Config{
			Hardware: &models.Hardware{
				AutomateNodeCount: 1,
				AutomateNodeIps:   []string{host},
			},
		}

		suc := NewSystemUserCheck(logger.NewLogrusStandardLogger(), port)
		ctr := suc.Run(config)

		require.Len(t, ctr, 1)
		require.NotNil(t, ctr[0].Result.Error)
		require.Equal(t, ctr[0].Result.Error.Code, http.StatusInternalServerError)
		assert.Equal(t, "error while parsing the response data:unexpected end of JSON input", ctr[0].Result.Error.Error())
	})

	t.Run("Nil Hardware", func(t *testing.T) {
		// Create a dummy server
		server, _, port := createDummyServer(t, http.StatusInternalServerError)
		defer server.Close()

		// Test data
		config := &models.Config{
			Hardware: nil,
		}

		suc := NewSystemUserCheck(logger.NewLogrusStandardLogger(), port)
		got := suc.Run(config)

		require.Len(t, got, 5)
		for _, v := range got {
			if v.CheckType == constants.BASTION {
				assert.Equal(t, constants.LOCALHOST, v.Host)
			}
			assert.Equal(t, constants.CHEF_INFRA_SERVER, got[1].NodeType)
			assert.Equal(t, constants.SYSTEM_USER, v.CheckType)
			assert.True(t, v.Result.Skipped)
		}
	})
}

// Helper function to create a dummy server
func createDummyServer(t *testing.T, requiredStatusCode int) (*httptest.Server, string, string) {
	if requiredStatusCode == http.StatusOK {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, constants.SYSTEM_USER_CHECK_API_PATH, r.URL.Path)
			reqParameters := r.URL.Query()
			assert.Equal(t, 0, len(reqParameters))

			if r.URL.Path == constants.SYSTEM_USER_CHECK_API_PATH {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(systemUser))
			}
		}))

		// Extract IP and port from the server's URL
		address := server.URL[strings.Index(server.URL, "//")+2:]
		colonIndex := strings.Index(address, ":")
		ip := address[:colonIndex]
		port := address[colonIndex+1:]

		return server, ip, port
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(requiredStatusCode)
	}))

	// Extract IP and port from the server's URL
	address := server.URL[strings.Index(server.URL, "//")+2:]
	colonIndex := strings.Index(address, ":")
	ip := address[:colonIndex]
	port := address[colonIndex+1:]

	return server, ip, port
}

func TestGetPortsForMockServer(t *testing.T) {
	fwc := NewSystemUserCheck(logger.NewLogrusStandardLogger(), "1234")
	resp := fwc.GetPortsForMockServer()

	assert.Equal(t, 0, len(resp))
}
