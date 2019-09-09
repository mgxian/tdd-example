package main

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{}

func (i *InMemoryUserStore) GetUserAPICallCount(user string) int {
	return 666
}

func main() {
	store := InMemoryUserStore{}
	server := &UserServer{&store}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
