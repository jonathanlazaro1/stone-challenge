package providers

import (
	"github.com/jonathanlazaro1/stone-challenge/domain/models"
)

// InvoiceProvider specifies which declarations Invoice repositories should implement
type InvoiceProvider interface {
	GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]models.Invoice, error)
}
