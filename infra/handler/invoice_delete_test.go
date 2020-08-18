package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestInvoiceDelete(t *testing.T) {
	json, _ := json.Marshal(makeInvoicePostModel())
	req, err := http.NewRequest("DELETE", makeParameterizedURL("1"), bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceDeleteHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusNoContent
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", "DELETE", endpoint, statusWant, statusGotten)
	}
}

func TestInvoiceDeleteWithUnparseableId(t *testing.T) {
	req, err := http.NewRequest("DELETE", makeParameterizedURL("notagreatId"), nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(InvoiceDeleteHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status code to be %v, but got %v", "DELETE", endpoint, statusWant, statusGotten)
	}

	bodyWant := errCouldntParseInvoiceID
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got \"%v\"", "DELETE", endpoint, bodyWant, bodyGotten)
	}
}
