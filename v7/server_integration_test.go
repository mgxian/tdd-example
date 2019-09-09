package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordAPICallsAndGetThem(t *testing.T) {
	store := NewInMemoryUserStore()
	server := UserServer{store}
	user := "will"

	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%s", user), nil)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)

	response := httptest.NewRecorder()
	request = newGetUserAPICallCountRequest(user)
	server.ServeHTTP(response, request)

	assertStatus(t, response.Code, http.StatusOK)
	assertCount(t, response.Body.String(), "3")
}
