package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestInvoiceUpdate(t *testing.T) {
	json, _ := json.Marshal(makeInvoicePostModel())
	req, err := http.NewRequest("PUT", makeParameterizedURL("1"), bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceUpdateHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusNoContent
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", "PUT", endpoint, statusWant, statusGotten)
	}
}

func TestInvoicePutWithFaultyProps(t *testing.T) {
	postModel := makeInvoicePostModel()
	postModel.Description = nil
	postModel.Amount = nil

	json, _ := json.Marshal(postModel)
	req, err := http.NewRequest("PUT", makeParameterizedURL("1"), bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceUpdateHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", "PUT", endpoint, statusWant, statusGotten)
	}
}

func TestInvoiceUpdateWithNoBody(t *testing.T) {
	json, _ := json.Marshal("")
	req, err := http.NewRequest("PUT", makeParameterizedURL("1"), bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceUpdateHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", "PUT", endpoint, statusWant, statusGotten)
	}

	bodyWant := errCouldntParsePostModel
	bodyGotten := rr.Body.String()

	if bodyWant != bodyGotten {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got value \"%v\"", "PUT", endpoint, bodyWant, bodyGotten)
	}
}

func TestInvoicePatch(t *testing.T) {
	postModel := makeInvoicePostModel()
	postModel.Description = nil
	postModel.Amount = nil
	json, _ := json.Marshal(postModel)
	req, err := http.NewRequest("PATCH", makeParameterizedURL("1"), bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceUpdateHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusNoContent
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", "PATCH", endpoint, statusWant, statusGotten)
	}
}
