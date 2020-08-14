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

// PutHandler handles a request to put an Invoice
func PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseInvoiceID)
		return
	}

	var model PostModel
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParsePostModel)
		return
	}

	model.id = id

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint(err))
		return
	}

	service := service.BuildInvoiceService()

	rowsAffected, err := service.Update(model.ToInvoice())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Couldn't update Invoice")
		return
	}

	if rowsAffected < 1 {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errInvoiceNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
