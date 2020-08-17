package pgsql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jonathanlazaro1/stone-challenge/config"
	// PGSQL driver
	_ "github.com/lib/pq"
)

// CreateConnection tries to create a DBConnection using PGSQL. It panics if is not possible to connect to the DB specified in .env
func CreateConnection() (*sql.DB, error) {
	cfg := config.GetConfig()
	var err error
	var db *sql.DB
	var connStr string

	if cfg.DBURL != "" {
		connStr = cfg.DBURL
	} else {
		connStr = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
			cfg.DBSSLMode)
	}

	db, err = sql.Open("postgres", connStr)
	if db == nil {
		return nil, err
	}

	for i := 1; i < 6; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		log.Printf("Connect to DB failed after 5 attempts: %v", err)
		return nil, err
	}
	return db, nil
}
