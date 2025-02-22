package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler serves the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Welcome to my Go API! xxx")
}
