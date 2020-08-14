package repository

import (
	domain "github.com/jonathanlazaro1/stone-challenge/domain/invoice"
)

// Invoice repository specifies which declarations Invoice repositories should implement
type Invoice interface {
	GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error)
	Get(id int) (*domain.Invoice, error)
	Add(invoice domain.Invoice) (int, error)
}
