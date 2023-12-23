package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// utils here
func MockHttpFunc(route string, method string, body map[string]string) (*http.Request, *httptest.ResponseRecorder, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}

	// Create a new mock request
	req, err := http.NewRequest(method, route, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, nil, err
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	return req, rr, nil
}
