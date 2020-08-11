package repository

import (
	"github.com/jonathanlazaro1/stone-challenge/domain/models"
)

// Invoice repository specifies which declarations Invoice repositories should implement
type Invoice interface {
	GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) (*[]models.Invoice, error)
}
