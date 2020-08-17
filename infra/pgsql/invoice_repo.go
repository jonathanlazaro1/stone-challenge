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

//TODO: change document filter to be "like" instead of "equal"
// GetMany fetches all invoices found on DB table invoice, according to the parameters given. It also returns the total count for the query made
func (repo *invoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]domain.Invoice, int64, error) {
	db := CreateConnection()
	defer db.Close()
	database := goqu.New("postgresql", db)

	goquWhere := goqu.Ex{"is_active": true}
	for k, v := range filterBy {
		goquWhere[k] = v
	}

	goquSQL := database.
		From("invoice").
		Where(goquWhere)

	// Counting rows
	count, err := goquSQL.Count()
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

	sql, _, _ := goquSQL.ToSQL()

	rows, err := database.Query(sql)

	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
		return nil, count, err
	}

	defer rows.Close()
	invoices := []domain.Invoice{}

	for rows.Next() {
		var invoice domain.Invoice
		err = rows.Scan(
			&invoice.ID,
			&invoice.ReferenceYear,
			&invoice.ReferenceMonth,
			&invoice.Document,
			&invoice.Description,
			&invoice.Amount,
			&invoice.IsActive,
			&invoice.CreatedAt,
			&invoice.DeactivatedAt)

		if err != nil {
			log.Printf("Unable to fetch invoices. %v", err)
			return nil, count, err
		}

		invoices = append(invoices, invoice)
	}
	return invoices, count, err
}

// Get finds an Invoice, given its Id
func (repo *invoiceRepository) Get(id int) (*domain.Invoice, error) {
	db := CreateConnection()
	defer db.Close()
	database := goqu.New("postgresql", db)

	goquWhere := goqu.Ex{
		"is_active": true,
		"id":        id,
	}

	goquSQL := database.From("invoice").Where(goquWhere)

	count, err := goquSQL.Count()
	if count < 1 {
		return nil, nil
	}

	sql, _, _ := goquSQL.ToSQL()
	var invoice domain.Invoice

	err = database.QueryRow(sql).Scan(
		&invoice.ID,
		&invoice.ReferenceYear,
		&invoice.ReferenceMonth,
		&invoice.Document,
		&invoice.Description,
		&invoice.Amount,
		&invoice.IsActive,
		&invoice.CreatedAt,
		&invoice.DeactivatedAt)

	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
		return nil, err
	}
	return &invoice, nil
}

// Add creates a new Invoice on DB
func (repo *invoiceRepository) Add(invoice domain.Invoice) (int, error) {
	db := CreateConnection()
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

	goquSQL := database.From("invoice").Insert().Rows(record).Returning("id")
	sql, _, _ := goquSQL.ToSQL()
	var id int

	err := database.QueryRow(sql).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

// Update updates an existent Invoice on DB
func (repo *invoiceRepository) Update(invoice domain.Invoice) (int64, error) {
	db := CreateConnection()
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
	sql, _, _ := goquSQL.ToSQL()

	var rowsAffected int64 = -1

	res, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return rowsAffected, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	return rowsAffected, err
}
