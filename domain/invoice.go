package domain

import (
	"database/sql"
	"encoding/json"
	"time"
)

// TODO: remove SQLNullTime things away from domain
// TODO: guarantee UTC time output

// JSONNullTime represents a time.Time type that can be null
type JSONNullTime struct {
	sql.NullTime
}

// MarshalJSON parses a NullableTimeType into a JSON (ISO 8601) type or null
func (v JSONNullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)
}

// Invoice represents an Invoice entity. A new Invoice must be obtained by using NewInvoice method
type Invoice struct {
	ID             int          `json:"id" db:"id"`
	ReferenceMonth int          `json:"referenceMonth" db:"reference_month"`
	ReferenceYear  int          `json:"referenceYear" db:"reference_year"`
	Document       string       `json:"document" db:"document"`
	Description    string       `json:"description" db:"description"`
	Amount         float64      `json:"amount" db:"amount" swaggertype:"number" format:"float"`
	IsActive       bool         `json:"isActive" db:"is_active"`
	CreatedAt      time.Time    `json:"createdAt" db:"created_at" swaggertype:"string" format:"date-time"`
	DeactivatedAt  JSONNullTime `json:"deactivatedAt" db:"deactivated_at" swaggertype:"string" format:"date-time"`
}

// NewInvoice creates a new instance of an Invoice
func NewInvoice() Invoice {
	i := &Invoice{}
	i.IsActive = true
	i.CreatedAt = time.Now().UTC()

	return *i
}

// Deactivate makes an Invoice inactive. It also sets the current UTC date/time of deactivation
func (invoice *Invoice) Deactivate() {
	invoice.IsActive = false
	invoice.DeactivatedAt.Time = time.Now().UTC()
}
