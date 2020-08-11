package test

import (
	"time"

	"github.com/brianvoe/gofakeit/v5"
	in "github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	rp "github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

func substr(str string, maxLen int) string {
	runes := []rune(str)
	if len(runes) < maxLen {
		return string(runes)
	}
	return string(runes[:maxLen-1])
}

type mockedInvoiceRepository struct {
	invoices []in.Invoice
}

func (mp *mockedInvoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]in.Invoice, error) {

	newInvoices := append(mp.invoices[:0], mp.invoices[:itemsPerPage]...)
	return newInvoices, nil
}

// MockInvoiceRepository creates a mocked implementation of an InvoiceRepository
func MockInvoiceRepository(size int) rp.Invoice {
	repo := &mockedInvoiceRepository{}
	for i := 0; i < size; i++ {
		repo.invoices = append(repo.invoices, MakeFakeInvoice())
	}

	return repo
}

// MakeFakeInvoice generates a pointer to an Invoice with fake populated data
func MakeFakeInvoice() in.Invoice {
	in := in.NewInvoice()
	in.ID = gofakeit.Number(1, 100000)

	currYear := time.Now().Year()
	in.ReferenceYear = gofakeit.Number(currYear-5, currYear-1)
	in.ReferenceMonth = gofakeit.Number(1, 12)

	in.Document = substr(gofakeit.Phrase(), 14)
	in.Description = substr(gofakeit.Phrase(), 256)
	in.Amount = gofakeit.Price(0.01, 1000000)

	return in
}
