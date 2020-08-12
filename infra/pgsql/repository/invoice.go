package repository

import (
	"log"

	"github.com/jonathanlazaro1/stone-challenge/domain/invoice"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
)

func getAllUsers() ([]invoice.Invoice, error) {
	db := pgsql.CreateConnection()
	defer db.Close()

	var invoices []invoice.Invoice

	sqlStatement := `SELECT () FROM invoices`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var invoice invoice.Invoice

		// unmarshal the row object to user
		// err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		invoices = append(invoices, invoice)
	}
	return invoices, err
}
