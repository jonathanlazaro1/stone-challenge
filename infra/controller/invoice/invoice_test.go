package invoice

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvoiceIndexHandler(t *testing.T) {
	verb := "GET"
	endpoint := "/api/v1/invoice"
	req, err := http.NewRequest(verb, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

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
