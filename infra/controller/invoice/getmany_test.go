package invoice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ==================== NO PARAMS ====================

func TestInvoiceGetManyWithoutParameters(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

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

// ==================== ITEMS PER PAGE ====================

func TestInvoiceGetManyWithItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?itemsperpage=20"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

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

func TestInvoiceGetManyWithUnparseableItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?itemsperpage=notagreatitemsperpage"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status code to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errParsingItemsPerPage
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}

func TestInvoiceGetManyHandlerWithExcessiveItemsPerPage(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice?itemsperpage=1445332"
	req, err := http.NewRequest(verb, endpoint, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status code to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := fmt.Sprintf(errMaxItemsPerPageAllowed, maxItemsPerPage)
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
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
	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

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
	handler := http.HandlerFunc(buildRequestFunction(GetManyHandler))

	handler.ServeHTTP(rr, req)

	statusWant := http.StatusBadRequest
	if statusGotten := rr.Code; statusGotten != statusWant {
		t.Errorf("%v %v failed: wanted status code to be %v, but got %v", verb, endpoint, statusWant, statusGotten)
	}

	bodyWant := errParsingPageNumber
	bodyGotten := rr.Body.String()
	if bodyGotten != bodyWant {
		t.Errorf("%v %v failed: wanted body to be \"%v\", but got \"%v\"", verb, endpoint, bodyWant, bodyGotten)
	}
}
