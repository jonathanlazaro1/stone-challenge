package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	verb := "GET"
	endpoint := "/"

	req, err := http.NewRequest(verb, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusOK
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := `Welcome! Here's a cup of ☕ for you`
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body \"%v\", but got  \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}
