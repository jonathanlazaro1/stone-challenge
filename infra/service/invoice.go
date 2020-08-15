package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	domain "github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/helpers"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql/repository"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice"
	uc "github.com/jonathanlazaro1/stone-challenge/usecase/invoice"
)

// PostModel represents an Invoice model to be persisted to an Invoice
type PostModel struct {
	ReferenceYear  *int     `json:"referenceYear,omitempty"`
	ReferenceMonth *int     `json:"referenceMonth,omitempty"`
	Document       *string  `json:"document,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Amount         *float64 `json:"amount,omitempty"`
}

// Validate verifies if a InvoicePostModel has valid data.
func (m PostModel) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ReferenceYear, validation.Required, validation.Min(1900), validation.Max(2100)),
		validation.Field(&m.ReferenceMonth, validation.Required, validation.Min(1), validation.Max(12)),
		validation.Field(&m.Document, validation.Required, validation.RuneLength(1, 14)),
		validation.Field(&m.Description, validation.Required, validation.RuneLength(1, 256)),
		validation.Field(&m.Amount, validation.Required, validation.Min(0.01)),
	)
}

// ToInvoice converts a PostModel to an Invoice instance
func (m PostModel) ToInvoice(invoiceToMerge *domain.Invoice) domain.Invoice {
	if invoiceToMerge == nil {
		invoice := domain.NewInvoice()
		invoice.ReferenceYear = *m.ReferenceYear
		invoice.ReferenceMonth = *m.ReferenceMonth
		invoice.Document = *m.Document
		invoice.Description = *m.Description
		invoice.Amount = *m.Amount

		return invoice
	}
	helpers.CopyIfNotNil(&m, invoiceToMerge)
	return *invoiceToMerge
}

// BuildInvoiceService builds a new InvoiceInteractor with the specified repository
func BuildInvoiceService() *invoice.Interactor {
	repo := repository.GetInvoiceRepository()
	service := uc.NewInteractor(repo)

	return service
}

// Put replaces Invoce values on the repository from the ones coming from the request.
// It ignores the server-control related ones, such as CreatedAt, IsActive and UpdatedAt
func Put(id int, postModel PostModel) (*domain.Invoice, error) {
	service := BuildInvoiceService()
	currentInvoice, err := service.Get(id)
	if err != nil {
		return nil, err
	} else if currentInvoice == nil {
		return nil, nil
	}
	updatedInvoice := postModel.ToInvoice(currentInvoice)

	rowCount, err := service.Update(updatedInvoice)
	if err != nil {
		return nil, err
	} else if rowCount < 1 {
		return nil, nil
	}

	currentInvoice, err = service.Get(id)
	if err != nil {
		return nil, err
	}
	return currentInvoice, nil
}

// Delete makes an Invoice inactive, which makes it unable to be retrieved/maintained
func Delete(id int) (bool, error) {
	service := BuildInvoiceService()
	invoice, err := service.Get(id)
	if err != nil {
		return false, err
	} else if invoice == nil {
		return false, nil
	}

	rowCount, err := service.Delete(*invoice)
	if err != nil {
		return false, err
	} else if rowCount < 1 {
		return false, nil
	}

	invoice, err = service.Get(id)
	if err != nil {
		return false, err
	}
	return invoice == nil, nil
}
