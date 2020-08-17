package invoice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const getByIDVerb = "GET"

func TestInvoiceGetById(t *testing.T) {
	getByIDEndpoint := fmt.Sprintf("%v/{id}", endpoint)
	req, err := http.NewRequest(getByIDVerb, makeParameterizedURL("1"), nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(GetHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusOK
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", getByIDVerb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", getByIDVerb, getByIDEndpoint, contentTypeWant, contentTypeGotten)
	}
}

func TestInvoiceGetByIdWithUnparseableId(t *testing.T) {
	req, err := http.NewRequest(getByIDVerb, makeParameterizedURL("notagreatId"), nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(makeParameterizedURL("{id}"), buildRequestFunction(GetHandler))
	router.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status code to be %v, but got %v", getByIDVerb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errCouldntParseInvoiceID
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got \"%v\"", getByIDVerb, endpoint, bodyWant, bodyGotten)
	}
}
