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
	env := config.GetConfig()

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("PGSQL Successfully connected!")
	return db
}
