package main

import (
	"fmt"
	"net/http"
)

func UserServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "6")
}
