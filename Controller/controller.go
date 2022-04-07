package controller

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s", name)
}
