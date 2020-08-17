package main

import (
	"github.com/jonathanlazaro1/stone-challenge/config"
	"github.com/jonathanlazaro1/stone-challenge/infra/pgsql"
)

func main() {
	config.Load()
	pgsql.Migrate()
}
