package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestBuildForwardRequest(t *testing.T) {
	bodyBytes := []byte("hello")
	targetUrl := "https://www.baidu.com/"
	apiKey := "secret"
	method := http.MethodPost

	request, err := buildForwardRequest(bodyBytes, method, targetUrl, apiKey)
	assert.NoError(t, err)

	assert.Equal(t, request.Method, method)
	assert.Equal(t, request.URL.String(), targetUrl)
	assert.Equal(t, request.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", apiKey))
	assert.Equal(t, request.Header.Get("Content-Type"), "application/json")
	requestBodyBytes, err := io.ReadAll(request.Body)
	assert.NoError(t, err)
	assert.Equal(t, requestBodyBytes, bodyBytes)
}
