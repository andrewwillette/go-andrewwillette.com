package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginPost(t *testing.T) {
	writer := httptest.NewRecorder()
	body := User {Username: "hello", Password: "passwordWorld"}
	request := httptest.NewRequest(http.MethodPost, loginEndpoint, body)
	loginPost(writer, request)
	assert.Equal(t, "", writer.Result().Status)
}