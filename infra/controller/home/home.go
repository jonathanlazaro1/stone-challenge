package home

import (
	"io"
	"net/http"
)

// IndexHandler greets the request caller with a cup of coffee
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `Welcome! Here's a cup of ☕ for you`)
}
