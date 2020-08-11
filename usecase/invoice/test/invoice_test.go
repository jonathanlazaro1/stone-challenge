package test

import (
	"testing"
)

func TestInvoiceGetMany(t *testing.T) {
	mockedRepo := MockInvoiceRepository(10)
	filterByMap := make(map[string]string)
	sortByMap := make(map[string]bool)

	want := 5

	invoices, err := mockedRepo.GetMany(want, 1, filterByMap, sortByMap)

	if err != nil {
		t.Errorf("GetMany failed: mocked repo generated an error. %v", err)
	}

	got := len(invoices)

	if got != want {
		t.Errorf("GetMany failed: wanted %v invoices, but got %v from mocked repo", want, got)
	}
}
