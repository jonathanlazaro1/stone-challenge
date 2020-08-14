package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// PostModel represents an Invoice model to be persisted to an Invoice
type PostModel struct {
	ReferenceYear  int     `json:"referenceYear"`
	ReferenceMonth int     `json:"referenceMonth"`
	Document       string  `json:"document"`
	Description    string  `json:"description"`
	Amount         float64 `json:"amount"`
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

// PostHandler handles a request to post an Invoice
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var model PostModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Couldn't parse invoice")
		return
	}

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
