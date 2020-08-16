package main

// @title Invoice API
// @version 1.0
// @description An API that takes care of Invoices

// @tag.name invoices
// @tag.description Invoice related operations
// @tag.name auth
// @tag.description Auth related operations

// @host localhost:8080
// @BasePath /api/v1

// @license.name MIT
// @license.url https://github.com/jonathanlazaro1/stone-challenge/blob/master/LICENSE

// @contact.name Jonathan Lazaro
// @contact.email jonathan.lazaro1@gmail.com

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization
// @tokenUrl /auth

import (
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/config"

	_ "github.com/jonathanlazaro1/stone-challenge/docs"

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
