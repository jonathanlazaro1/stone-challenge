package pgsql

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// execSQL executes an SQL statement against a DB connection
func execSQL(db *sql.DB, sql string) error {
	_, err := db.Exec(sql)
	return err
}

// Migrate tries to update the PGSQL database that is set on environment, taking it to the state needed by application to work with
func Migrate() error {
	db, err := CreateConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	var sqlFiles []string
	migrationsDir := "./infra/pgsql/migrations/"

	// Get all .sql files on migrations folder
	err = filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".sql" {
			return nil
		}
		sqlFiles = append(sqlFiles, path)
		return nil
	})
	if err != nil {
		return err
	}

	// Go through the .sql files found on folder to exec the SQL statements
	for _, file := range sqlFiles {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		sql := string(bytes)

		err = execSQL(db, sql)
		if err != nil {
			return err
		}
	}
	log.Println("PGSQL migrations applied successfully")
	return nil
}
