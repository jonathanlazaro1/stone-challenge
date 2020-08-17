package main

import (
	"log"
	"os"

	"github.com/jonathanlazaro1/stone-challenge/config"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
)

func main() {
	config.Load()
	err := pgsql.Migrate()
	if err != nil {
		log.Fatalf("Could not execute migrations: %v", err)
	}

	os.Exit(0)
}
