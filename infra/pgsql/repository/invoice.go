package repository

import (
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"

	// Goqu PGSQL dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type invoiceRepository struct {
}

// GetInvoiceRepository returns an implementation of InvoiceRepository that relies on a PGSQL DB
func GetInvoiceRepository() repository.Invoice {
	return &invoiceRepository{}
}

// GetMany fetches all invoices found on DB table invoice, according to the parameters given. It also returns the total count for the query made
func (repo *invoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]invoice.Invoice, int64, error) {
	db := pgsql.CreateConnection()
	defer db.Close()
	database := goqu.New("postgresql", db)

	goquWhere := goqu.Ex{
		"is_active": true,
	}
	for k, v := range filterBy {
		goquWhere[k] = v
	}

	goquSQL := database.
		From("invoice").
		Select(
			goqu.C("id"),
			goqu.C("reference_year"),
			goqu.C("reference_month"),
			goqu.C("document"),
			goqu.C("description"),
			goqu.C("amount"),
			goqu.C("is_active"),
			goqu.C("created_at"),
			goqu.C("deactivated_at")).
		Where(goquWhere)

	// Counting rows
	count, err := goquSQL.Count()
	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
		return nil, count, err
	}

	goquSQL = goquSQL.
		Limit(uint(itemsPerPage)).
		Offset(uint(itemsPerPage * (page - 1)))

	for k, descend := range sortBy {
		if descend {
			goquSQL = goquSQL.OrderAppend(goqu.I(k).Desc())
		} else {
			goquSQL = goquSQL.OrderAppend(goqu.I(k).Asc())
		}
	}

	sql, _, _ := goquSQL.ToSQL()
	fmt.Println(sql)

	rows, err := database.Query(sql)

	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
		return nil, count, err
	}

	defer rows.Close()
	invoices := []invoice.Invoice{}

	for rows.Next() {
		var invoice invoice.Invoice
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
