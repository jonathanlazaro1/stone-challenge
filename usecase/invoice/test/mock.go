package test

import (
	"time"

	"github.com/jonathanlazaro1/stone-challenge/helpers"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/jonathanlazaro1/stone-challenge/domain"
	rp "github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

type mockedInvoiceRepository struct {
	invoices []domain.Invoice
}

func (mp *mockedInvoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	newInvoices := append(mp.invoices[:0], mp.invoices[:itemsPerPage]...)
	return newInvoices, int64(len(mp.invoices)), nil
}

func (mp *mockedInvoiceRepository) Get(id int) (*domain.Invoice, error) {
	return &mp.invoices[0], nil
}

func (mp *mockedInvoiceRepository) Add(invoice domain.Invoice) (int, error) {
	mp.invoices = append(mp.invoices, invoice)
	return len(mp.invoices), nil
}

func (mp *mockedInvoiceRepository) Update(invoice domain.Invoice) (int64, error) {
	mp.invoices = append(mp.invoices, invoice)
	return int64(len(mp.invoices)), nil
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
func MakeFakeInvoice() domain.Invoice {
	in := domain.NewInvoice()
	in.ID = gofakeit.Number(1, 100000)

	currYear := time.Now().Year()
	in.ReferenceYear = gofakeit.Number(currYear-5, currYear-1)
	in.ReferenceMonth = gofakeit.Number(1, 12)

	in.Document = helpers.Substring(gofakeit.Phrase(), 14)
	in.Description = helpers.Substring(gofakeit.Phrase(), 256)
	in.Amount = gofakeit.Price(0.01, 1000000)

	return in
}
