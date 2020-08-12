package invoice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TODO: improve tests to check if passed parameteres are coming to handler as expected
//TODO: do tests pass with a sensitive case param?

// ==================== NO PARAMS ====================
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
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}

// ==================== PAGE NUMBER ====================

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
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}

func TestInvoiceGetManyHandlerWithUnparseablePageNumber(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?p=impossible"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errParsingPageNumber
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}

// ==================== ITEMS PER PAGE ====================

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
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	contentTypeWant := `application/json`
	contentTypeGotten := rr.Header().Get("Content-Type")
	if contentTypeGotten != contentTypeWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got  \"%v\"", verb, endpoint, contentTypeWant, contentTypeGotten)
	}
}

func TestInvoiceGetManyHandlerWithUnparseableItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?ipp=notagreatitemsperpage"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errParsingItemsPerPage
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}

func TestInvoiceGetManyHandlerWithExcessiveItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?ipp=1445332"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetManyHandler)

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted body to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := fmt.Sprintf(errMaxItemsPerPageAllowed, maxItemsPerPage)
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted Content-Type to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}
