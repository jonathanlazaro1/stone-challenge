package router

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Router returns a instance of mux.Router ready to listen and respond to requests
func Router() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json")))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusTemporaryRedirect)
	})

	apiv1 := router.PathPrefix("/api/v1/").Subrouter()

	addAuthHandler(apiv1.PathPrefix("/auth").Subrouter())
	addInvoiceHandler(apiv1.PathPrefix("/invoice").Subrouter())

	return router
}
