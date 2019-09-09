package main

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct {
	store map[string]int
}

func (i *InMemoryUserStore) GetUserAPICallCount(user string) int {
	return i.store[user]
}

func (i *InMemoryUserStore) RecordAPICall(user string) {
	i.store[user]++
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		store: make(map[string]int),
	}
}

func main() {
	store := NewInMemoryUserStore()
	server := &UserServer{store}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
