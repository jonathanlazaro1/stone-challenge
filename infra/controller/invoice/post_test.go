package invoice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const verb = "POST"
const endpoint = "/api/v1/invoice"

func makeInvoicePostModel() service.PostModel {
	description := "The old description"
	refYear := 2020
	refMonth := 8
	doc := "current doc"
	amount := 2.42

	model := service.PostModel{}
	model.ReferenceYear = &refYear
	model.ReferenceMonth = &refMonth
	model.Document = &doc
	model.Description = &description
	model.Amount = &amount

	return model
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
