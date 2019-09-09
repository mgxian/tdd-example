package main

import (
	"fmt"
	"net/http"
)

type UserStore interface {
	GetUserAPICallCount(user string) int
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.showAPICallCount(w, r)
	case http.MethodPost:
		u.processAPICall(w, r)
	}
}

func (u *UserServer) showAPICallCount(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[len("/users/"):]
	apiCallCount := u.store.GetUserAPICallCount(user)
	if apiCallCount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, apiCallCount)
}

func (u *UserServer) processAPICall(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
