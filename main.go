package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
