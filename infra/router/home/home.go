package home

import "net/http"

// Index greets the request caller with a cup of coffee
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Welcome! Here's a cup of ☕ for you`))
}
