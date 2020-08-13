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
func Migrate() {
	db := CreateConnection()
	defer db.Close()

	var sqlFiles []string
	migrationsDir := "./infra/pgsql/migrations/"

	// Get all .sql files on migrations folder
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".sql" {
			return nil
		}
		sqlFiles = append(sqlFiles, path)
		return nil
	})
	if err != nil {
		log.Fatalln("Could not execute PGSQL DB migration")
	}

	// Go through the .sql files found on folder to exec the SQL statements
	for _, file := range sqlFiles {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Could not read file %v", file)
		}
		sql := string(bytes)

		err = execSQL(db, sql)
		if err != nil {
			log.Fatalf("Could not execute migration %v: %v", file, err)
		}
	}
	log.Println("Migrations applied successfully")
}
