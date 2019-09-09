package main

import (
	"fmt"
	"net/http"
)

type UserStore interface {
	GetUserAPICallCount(user string) int
	RecordAPICall(user string)
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[len("/users/"):]
	switch r.Method {
	case http.MethodGet:
		u.showAPICallCount(w, user)
	case http.MethodPost:
		u.processAPICall(w, user)
	}
}

func (u *UserServer) showAPICallCount(w http.ResponseWriter, user string) {
	apiCallCount := u.store.GetUserAPICallCount(user)
	if apiCallCount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, apiCallCount)
}

func (u *UserServer) processAPICall(w http.ResponseWriter, user string) {
	u.store.RecordAPICall(user)
	w.WriteHeader(http.StatusAccepted)
}
