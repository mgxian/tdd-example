package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	t.Run("return will's api call count", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/will", nil)
		response := httptest.NewRecorder()
		UserServer(response, request)
		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
