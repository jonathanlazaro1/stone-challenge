package invoice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const verb = "POST"
const endpoint = "/api/v1/invoice"

func makeInvoicePostModel() PostModel {
	return PostModel{
		ReferenceMonth: 8,
		ReferenceYear:  2020,
		Document:       "202008001",
		Description:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus vitae tortor mi. Fusce lobortis sed est eu sollicitudin. Vestibulum et eleifend neque, sed laoreet magna. Suspendisse ut molestie nisl.",
		Amount:         1.45,
	}
}

func TestInvoicePost(t *testing.T) {

	json, _ := json.Marshal(makeInvoicePostModel())
	req, err := http.NewRequest(verb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusCreated
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyGotten := rr.Body.String()
	_, err = strconv.Atoi(bodyGotten)

	if err != nil {
		t.Errorf("%v %v failed: wanted body to be an integer, but got value \"%v\"", verb, endpoint, bodyGotten)
	}
}

func TestInvoicePostWithFaultyProps(t *testing.T) {
	emptyBody := struct{}{}
	json, _ := json.Marshal(emptyBody)
	req, err := http.NewRequest(verb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}
}

func TestInvoicePostWithNoBody(t *testing.T) {
	json, _ := json.Marshal("")
	req, err := http.NewRequest(verb, endpoint, bytes.NewBuffer(json))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errCouldntParsePostModel
	bodyGotten := rr.Body.String()

	if bodyWant != bodyGotten {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got value \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}
