package invoice

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//TODO: improve tests to check if passed parameteres are coming to handler as expected

func TestInvoiceGetManyHandlerWithoutParameters(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusOK
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted body to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}

func TestInvoiceGetManyHandlerWithPageNumber(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?p=2"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusOK
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted body to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}

func TestInvoiceGetManyHandlerWithItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?ipp=20"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusOK
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted body to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}
