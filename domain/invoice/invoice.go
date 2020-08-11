package invoice

import (
	"time"
)

// Invoice represents an Invoice entity. A new Invoice must be obtained by using NewInvoice method
type Invoice struct {
	ID             int
	ReferenceMonth int
	ReferenceYear  int
	Document       string
	Description    string
	Amount         float64
	IsActive       bool
	CreatedAt      time.Time
	DeactivatedAt  time.Time
}

// NewInvoice creates a new instance of an Invoice
func NewInvoice() Invoice {
	i := &Invoice{}
	i.IsActive = true
	i.CreatedAt = time.Now().UTC()

	return *i
}
