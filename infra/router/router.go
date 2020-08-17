package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Router returns a instance of mux.Router ready to listen and respond to requests
func Router() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.PathPrefix("/swagger").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("https://stone-invoice-api.herokuapp.com/swagger/doc.json")))

	apiv1 := router.PathPrefix("/api/v1/").Subrouter()

	addAuthHandler(apiv1.PathPrefix("/auth").Subrouter())
	addInvoiceHandler(apiv1.PathPrefix("/invoice").Subrouter())

	return router
}
