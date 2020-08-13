package pgsql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jonathanlazaro1/stone-challenge/config"
	// PGSQL driver
	_ "github.com/lib/pq"
)

// CreateConnection tries to create a DBConnection using PGSQL. It panics if is not possible to connect to the DB specified in .env
func CreateConnection() *sql.DB {
	cfg := config.GetConfig()

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
