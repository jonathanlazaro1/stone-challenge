package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const postVerb = "POST"

func TestInvoicePost(t *testing.T) {

	json, _ := json.Marshal(makeInvoicePostModel())
	req, err := http.NewRequest(postVerb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(InvoicePostHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusCreated
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", postVerb, endpoint, statusWant, statusGotten)
	}

	bodyGotten := rr.Body.String()
	_, err = strconv.Atoi(bodyGotten)

	if err != nil {
		t.Errorf("%v %v failed: wanted body to be an integer, but got value \"%v\"", postVerb, endpoint, bodyGotten)
	}
}

func TestInvoicePostWithFaultyProps(t *testing.T) {
	emptyBody := struct{}{}
	json, _ := json.Marshal(emptyBody)
	req, err := http.NewRequest(postVerb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(InvoicePostHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", postVerb, endpoint, statusWant, statusGotten)
	}
}

func TestInvoicePostWithNoBody(t *testing.T) {
	json, _ := json.Marshal("")
	req, err := http.NewRequest(postVerb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(InvoicePostHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", postVerb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errCouldntParsePostModel
	bodyGotten := rr.Body.String()

	if bodyWant != bodyGotten {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got value \"%v\"", postVerb, endpoint, bodyWant, bodyGotten)
	}
}
