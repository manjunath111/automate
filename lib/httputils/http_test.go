package httputils_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chef/automate/lib/httputils"
	"github.com/chef/automate/lib/logger"
	"github.com/stretchr/testify/assert"
)

func TestMakeRequestSuccessfulRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var requestBody map[string]interface{}
		err = json.Unmarshal(body, &requestBody)
		assert.NoError(t, err)

		expectedBody := map[string]interface{}{
			"key1": "value1",
			"key2": "42",
		}
		assert.Equal(t, expectedBody, requestBody)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}))
	defer server.Close()

	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	url := server.URL
	requestBody := map[string]interface{}{
		"key1": "value1",
		"key2": "42",
	}
	resp, responseBody, err := client.MakeRequest(http.MethodPost, url, requestBody)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expectedResponseBody := "Success"
	assert.Equal(t, expectedResponseBody, string(responseBody))
}

func TestMakeRequestRequestBodyError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	resp, responseBody, err := client.MakeRequest(http.MethodPost, "https://example.com", make(chan int))

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "failed to marshal request body")
}

func TestMakeRequestRequestError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	resp, responseBody, err := client.MakeRequest(http.MethodGet, "invalid-url", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "failed to make HTTP request")
}

func TestMakeRequestConnectionError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)

	client := httputils.NewClientWithTimeout(1*time.Millisecond, logger)

	resp, responseBody, err := client.MakeRequest(http.MethodGet, "http://non-existent-server", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestMakeRequestWithHeaderSuccessfulRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var requestBody map[string]interface{}
		err = json.Unmarshal(body, &requestBody)
		assert.NoError(t, err)

		expectedBody := map[string]interface{}{
			"key1": "value1",
			"key2": "42",
		}
		assert.Equal(t, expectedBody, requestBody)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}))
	defer server.Close()

	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	url := server.URL
	requestBody := map[string]interface{}{
		"key1": "value1",
		"key2": "42",
	}
	resp, responseBody, err := client.MakeRequestWithHeaders(http.MethodPost, url, requestBody, "Content-Type", "application/json")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expectedResponseBody := "Success"
	assert.Equal(t, expectedResponseBody, string(responseBody))
}

func TestMakeRequestWithHeaderConnectionError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)

	client := httputils.NewClientWithTimeout(1*time.Millisecond, logger)

	resp, responseBody, err := client.MakeRequestWithHeaders(http.MethodGet, "http://non-existent-server", nil, "Content-Type", "application/json")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestMakeRequestWithHeaderRequestError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	resp, responseBody, err := client.MakeRequestWithHeaders(http.MethodGet, "invalid-url", nil, "Content-Type", "application/json")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "failed to make HTTP request")
}

func TestMakeRequestWithHeaderRequestBodyError(t *testing.T) {
	logger, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	client := httputils.NewClient(logger)

	resp, responseBody, err := client.MakeRequestWithHeaders(http.MethodPost, "https://example.com", make(chan int), "Content-Type", "application/json")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responseBody)
	assert.Contains(t, err.Error(), "failed to marshal request body")
}
