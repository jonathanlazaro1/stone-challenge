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

func updateHandler(w http.ResponseWriter, r *http.Request) {
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

	newInvoice, err := service.Put(id, model)
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

// PutHandler handles a request to put an Invoice
func PutHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r)
}

// PatchHandler handles a request to patch an Invoice
func PatchHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r)
}
