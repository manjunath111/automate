package trigger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chef/automate/lib/logger"
	"github.com/gofiber/fiber/v2"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
)

func RunCheck(config models.Config, log logger.Logger, port string, path string, depState string) []models.CheckTriggerResponse {
	var result []models.CheckTriggerResponse
	count := config.Hardware.AutomateNodeCount +
		config.Hardware.ChefInfraServerNodeCount +
		config.Hardware.PostgresqlNodeCount +
		config.Hardware.OpenSearchNodeCount

	outputCh := make(chan models.CheckTriggerResponse)

	// added one for bastion node
	if path == constants.SOFTWARE_VERSION_CHECK_API_PATH || path == constants.SYSTEM_RESOURCE_CHECK_API_PATH {
		count = count + 1
		endpoint := prepareEndpoint(path, "127.0.0.1", port, constants.BASTION, depState)
		go TriggerCheckAPI(endpoint, "127.0.0.1", constants.BASTION, http.MethodGet, outputCh, nil)
	}

	for _, ip := range config.Hardware.AutomateNodeIps {
		endpoint := prepareEndpoint(path, ip, port, constants.AUTOMATE, depState)
		go TriggerCheckAPI(endpoint, ip, constants.AUTOMATE, http.MethodGet, outputCh, nil)
	}
	for _, ip := range config.Hardware.ChefInfraServerNodeIps {
		endpoint := prepareEndpoint(path, ip, port, constants.CHEF_INFRA_SERVER, depState)
		go TriggerCheckAPI(endpoint, ip, constants.CHEF_INFRA_SERVER, http.MethodGet, outputCh, nil)
	}
	for _, ip := range config.Hardware.OpenSearchNodeIps {
		endpoint := prepareEndpoint(path, ip, port, constants.OPENSEARCH, depState)
		go TriggerCheckAPI(endpoint, ip, constants.OPENSEARCH, http.MethodGet, outputCh, nil)
	}
	for _, ip := range config.Hardware.PostgresqlNodeIps {
		endpoint := prepareEndpoint(path, ip, port, constants.POSTGRESQL, depState)
		go TriggerCheckAPI(endpoint, ip, constants.POSTGRESQL, http.MethodGet, outputCh, nil)
	}

	for i := 0; i < count; i++ {
		select {
		case res := <-outputCh:
			result = append(result, res)
		}
	}

	return result
}

func prepareEndpoint(path, ip, port, nodeType, depState string) string {
	endPoint := ""
	if path == constants.SOFTWARE_VERSION_CHECK_API_PATH {
		endPoint = fmt.Sprintf("http://%s:%s%s?node_type=%s", ip, port, path, nodeType)

	} else if path == constants.SYSTEM_RESOURCE_CHECK_API_PATH {
		endPoint = fmt.Sprintf("http://%s:%s%s?node_type=%s&deployment_state=%s", ip, port, path, nodeType, depState)

	} else if path == constants.SYSTEM_USER_CHECK_API_PATH {
		endPoint = fmt.Sprintf("http://%s:%s%s", ip, port, path)
	}

	return endPoint
}

// TriggerCheckAPI triggers the API with provided enpoint,host,nodetype,method and reqbody.
// EndPoint should be string
// The method can be GET,POST,PUT
// Return an channel for output
// In case of error, error is stored in output channel itself
func TriggerCheckAPI(endPoint, host, nodeType, method string, output chan<- models.CheckTriggerResponse, reqBody interface{}) {
	var ctr models.CheckTriggerResponse

	reader, err := interfaceToIOReader(reqBody)
	if err != nil {
		output <- models.CheckTriggerResponse{
			Host:     host,
			NodeType: nodeType,
			Result: models.ApiResult{
				Passed: false,
				Error: &fiber.Error{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("error while reading the request body: %s", err.Error()),
				},
			},
		}
		return
	}

	req, err := http.NewRequest(method, endPoint, reader)
	if err != nil {
		output <- models.CheckTriggerResponse{
			Host:     host,
			NodeType: nodeType,
			Result: models.ApiResult{
				Passed: false,
				Error: &fiber.Error{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("error while creating the request:%s", err.Error()),
				},
			},
		}
		return
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		output <- models.CheckTriggerResponse{
			Host:     host,
			NodeType: nodeType,
			Result: models.ApiResult{
				Passed: false,
				Error: &fiber.Error{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("error while connecting to the endpoint:%s", err.Error()),
				},
			},
		}
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		output <- models.CheckTriggerResponse{
			Host:     host,
			NodeType: nodeType,
			Result: models.ApiResult{
				Passed: false,
				Error: &fiber.Error{
					Code:    resp.StatusCode,
					Message: "error while connecting to the endpoint, received invalid status code",
				},
			},
		}
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&ctr); err != nil {
		output <- models.CheckTriggerResponse{
			Host:     host,
			NodeType: nodeType,
			Result: models.ApiResult{
				Passed: false,
				Error: &fiber.Error{
					Code:    http.StatusInternalServerError,
					Message: fmt.Sprintf("error while parsing the response data:%s", err.Error()),
				},
			},
		}
		return
	}

	ctr.Host = host
	ctr.NodeType = nodeType
	output <- ctr
}

func interfaceToIOReader(body interface{}) (io.Reader, error) {
	var reader io.Reader
	if body != nil {
		bx, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		reader = bytes.NewBuffer(bx)

	}
	return reader, nil
}