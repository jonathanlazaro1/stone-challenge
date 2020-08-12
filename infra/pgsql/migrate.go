package pgsql

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Migrate tries to update the PGSQL database that is set on environment, to put it on the DB state needed by application to work with
func Migrate() {
	var files []string
	migrationsDir := "./migrations/"
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatalln("Could not execute migration")
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
