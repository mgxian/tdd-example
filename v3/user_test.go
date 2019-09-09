package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	apiCallCounts map[string]int
}

func (s *StubUserStore) GetUserAPICallCount(user string) int {
	return s.apiCallCounts[user]
}

func TestGetUsers(t *testing.T) {
	store := StubUserStore{
		apiCallCounts: map[string]int{
			"will":   6,
			"mgxian": 8,
		},
	}
	server := &UserServer{&store}
	t.Run("return will's api call count", func(t *testing.T) {
		user := "will"
		request := newGetUserAPICallCountRequest(user)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertCount(t, response.Body.String(), "6")
	})

	t.Run("return mgxian's api call count", func(t *testing.T) {
		user := "mgxian"
		request := newGetUserAPICallCountRequest(user)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertCount(t, response.Body.String(), "8")
	})
}

func newGetUserAPICallCountRequest(user string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", user), nil)
	return request
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong status code got %d, want %d", got, want)
	}
}

func assertCount(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got % q, want % q", got, want)
	}
}
