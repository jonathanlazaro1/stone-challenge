package models

import (
	"time"
)

// Invoice represents an Invoice entity
type Invoice struct {
	ID             int
	ReferenceMonth int
	Document       string
	Description    string
	Amount         float32
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
