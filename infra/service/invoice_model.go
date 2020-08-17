package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jonathanlazaro1/stone-challenge/domain"
	"github.com/jonathanlazaro1/stone-challenge/helpers"
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
