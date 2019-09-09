package main

import (
	"fmt"
	"net/http"
)

func UserServer(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[len("/users/"):]
	if user == "will" {
		fmt.Fprint(w, "6")
		return
	}

	if user == "mgxian" {
		fmt.Fprint(w, "8")
		return
	}
}
