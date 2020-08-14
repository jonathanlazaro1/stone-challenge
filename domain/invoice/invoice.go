package invoice

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
	ID             int          `json:"id"`
	ReferenceMonth int          `json:"referenceMonth"`
	ReferenceYear  int          `json:"referenceYear"`
	Document       string       `json:"document"`
	Description    string       `json:"description"`
	Amount         float64      `json:"amount"`
	IsActive       bool         `json:"isActive"`
	CreatedAt      time.Time    `json:"createdAt"`
	DeactivatedAt  JSONNullTime `json:"deactivatedAt"`
}

// NewInvoice creates a new instance of an Invoice
func NewInvoice() Invoice {
	i := &Invoice{}
	i.IsActive = true
	i.CreatedAt = time.Now().UTC()

	return *i
}
