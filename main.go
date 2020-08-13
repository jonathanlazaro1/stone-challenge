package main

import (
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/config"

	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
	"github.com/jonathanlazaro1/stone-challenge/infra/router"
)

func main() {
	config.Load()

	pgsql.Migrate()
	r := router.Router()

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
