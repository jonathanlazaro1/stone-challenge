package repository

import (
	"fmt"
	"log"

	"github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
	"github.com/jonathanlazaro1/stone-challenge/usecase/invoice/repository"
)

type invoiceRepository struct {
}

// GetInvoiceRepository returns an implementation of InvoiceRepository that relies on a PGSQL DB
func GetInvoiceRepository() repository.Invoice {
	return &invoiceRepository{}
}

// GetMany fetches all invoices found on DB table invoice, according to the parameters given
func (repo *invoiceRepository) GetMany(itemsPerPage int, page int, filterBy map[string]string, sortBy map[string]bool) ([]invoice.Invoice, error) {
	db := pgsql.CreateConnection()
	defer db.Close()

	invoices := []invoice.Invoice{}

	sqlStatement := "SELECT id, reference_year, reference_month, document, description, amount, is_active, created_at, deactivated_at FROM invoice"
	sqlWhere := "WHERE is_active = true"
	sqlOrderBy := ""
	sqlLimitOffset := fmt.Sprintf("LIMIT %v OFFSET %v;", itemsPerPage, (itemsPerPage * (page - 1)))

	if len(filterBy) > 0 {
		for k, v := range filterBy {
			sqlWhere = fmt.Sprintf("%v AND %v = %v", sqlWhere, k, v)
		}
	}

	// if len(orderBy) > 0 {
	// 	sqlOrderBy = "ORDER BY"
	// 	for k, v := range filterBy {
	// 		sqlOrderBy = fmt.Sprintf("%v AND %v = %v", sqlWhere, k, v)
	// 	}
	// }

	rows, err := db.Query(fmt.Sprintf("%v %v %v %v;", sqlStatement, sqlWhere, sqlOrderBy, sqlLimitOffset))

	if err != nil {
		log.Printf("Unable to fetch invoices. %v", err)
	}

	defer rows.Close()

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
			return nil, err
		}

		invoices = append(invoices, invoice)
	}
	return invoices, err
}
