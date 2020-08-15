package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errCouldntParsePostModel = "Couldn't parse invoice"

// PostHandler handles a request to post an Invoice
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var model service.PostModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParsePostModel)
		return
	}

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint(err))
		return
	}

	service := service.BuildInvoiceService()

	id, err := service.Add(model.ToInvoice(nil))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Couldn't create Invoice")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, strconv.Itoa(id))
}
