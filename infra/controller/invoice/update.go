package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// UpdateHandler handles a request to update an Invoice
func UpdateHandler(svc service.Invoice) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errCouldntParseInvoiceID)
			return
		}

		var model service.PostModel
		err = json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errCouldntParsePostModel)
			return
		}

		if r.Method == "PUT" {
			err = model.Validate()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, fmt.Sprint(err))
				return
			}
		}

		svc := service.BuildInvoiceService()
		newInvoice, err := svc.Update(id, model)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Couldn't update Invoice")
			return
		} else if newInvoice == nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, errInvoiceNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
