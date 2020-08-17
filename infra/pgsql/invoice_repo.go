package pgsql

import (
	"fmt"
	"log"

	"github.com/jonathanlazaro1/stone-challenge/usecase"

	"github.com/doug-martin/goqu/v9"
	"github.com/jonathanlazaro1/stone-challenge/domain"

	// Goqu PGSQL dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type invoiceRepository struct {
}

// GetInvoiceRepository returns an implementation of InvoiceRepository that relies on a PGSQL DB
func GetInvoiceRepository() usecase.InvoiceRepository {
	return &invoiceRepository{}
}

// GetMany fetches all invoices found on DB table invoice, according to the parameters given. It also returns the total count for the query made
func (repo *invoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	var count int64 = -1
	db, err := CreateConnection()
	if err != nil {
		return nil, count, err
	}
	defer db.Close()
	database := goqu.New("postgresql", db)

	goquSQL := database.
		From("invoice").
		Where(goqu.C("is_active").Eq(true))

	for k, v := range filterBy {
		if k == "document" {
			goquSQL = goquSQL.Where(goqu.C(k).Like(fmt.Sprintf("%c%v%c", '\u0025', v, '\u0025')))
		} else {
			goquSQL = goquSQL.Where(goqu.C(k).Eq(v))
		}
	}

	// Counting rows
	count, err = goquSQL.Count()
	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
		return nil, count, err
	}

	goquSQL = goquSQL.Limit(uint(itemsPerPage)).Offset(uint(itemsPerPage * (page - 1)))

	for k, descend := range sortBy {
		if descend {
			goquSQL = goquSQL.OrderAppend(goqu.I(k).Desc())
		} else {
			goquSQL = goquSQL.OrderAppend(goqu.I(k).Asc())
		}
	}

	invoices := []domain.Invoice{}
	if err := goquSQL.ScanStructs(&invoices); err != nil {
		return nil, count, err
	}
	return invoices, count, nil
}

// Get finds an Invoice, given its Id
func (repo *invoiceRepository) Get(id int) (*domain.Invoice, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	database := goqu.New("postgresql", db)

	goquSQL := database.
		From("invoice").
		Where(
			goqu.C("is_active").Eq(true),
			goqu.C("id").Eq(id))

	count, err := goquSQL.Count()
	if count < 1 {
		return nil, nil
	}

	var invoice domain.Invoice
	rowWasFound, err := goquSQL.ScanStruct(&invoice)
	if err != nil {
		return nil, err
	}
	if !rowWasFound {
		return nil, nil
	}

	return &invoice, nil
}

// Add creates a new Invoice on DB
func (repo *invoiceRepository) Add(invoice domain.Invoice) (int, error) {
	var id int
	db, err := CreateConnection()
	if err != nil {
		return id, err
	}
	defer db.Close()
	database := goqu.New("postgresql", db)

	record := goqu.Record{
		"reference_year":  invoice.ReferenceYear,
		"reference_month": invoice.ReferenceMonth,
		"document":        invoice.Document,
		"description":     invoice.Description,
		"amount":          invoice.Amount,
		"is_active":       invoice.IsActive,
		"created_at":      invoice.CreatedAt,
		"deactivated_at":  invoice.DeactivatedAt,
	}

	_, err = database.From("invoice").Insert().Rows(record).Returning("id").Executor().ScanVal(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

// Update updates an existent Invoice on DB
func (repo *invoiceRepository) Update(invoice domain.Invoice) (int64, error) {
	var rowsAffected int64 = -1
	db, err := CreateConnection()
	if err != nil {
		return rowsAffected, err
	}
	defer db.Close()
	database := goqu.New("postgresql", db)

	record := goqu.Record{
		"reference_year":  invoice.ReferenceYear,
		"reference_month": invoice.ReferenceMonth,
		"document":        invoice.Document,
		"description":     invoice.Description,
		"amount":          invoice.Amount,
		"is_active":       invoice.IsActive,
		"created_at":      invoice.CreatedAt,
		"deactivated_at":  invoice.DeactivatedAt,
	}

	goquSQL := database.From("invoice").
		Where(goqu.Ex{"id": invoice.ID}).
		Update().
		Set(record)

	res, err := goquSQL.Executor().Exec()
	if err != nil {
		return rowsAffected, err
	}

	rowsAffected, err = res.RowsAffected()

	return rowsAffected, err
}
