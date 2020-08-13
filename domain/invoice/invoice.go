package invoice

import (
	"database/sql"
	"encoding/json"
	"time"
)

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
	ID             int
	ReferenceMonth int
	ReferenceYear  int
	Document       string
	Description    string
	Amount         float64
	IsActive       bool
	CreatedAt      time.Time
	DeactivatedAt  JSONNullTime
}

// NewInvoice creates a new instance of an Invoice
func NewInvoice() Invoice {
	i := &Invoice{}
	i.IsActive = true
	i.CreatedAt = time.Now().UTC()

	return *i
}
