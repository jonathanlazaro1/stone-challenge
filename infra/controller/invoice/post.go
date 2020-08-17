package invoice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

const errCouldntParsePostModel = "Couldn't parse invoice"

// PostHandler handles a request to post an Invoice
func PostHandler(svc service.InvoiceService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		id, err := svc.Add(model.ToInvoice(nil))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Couldn't create Invoice")
			return
		}

		w.Header().Add("Location", fmt.Sprintf("%v%v%v/%v", r.URL.Scheme, r.Host, r.RequestURI, id))
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, strconv.Itoa(id))
	}
}
