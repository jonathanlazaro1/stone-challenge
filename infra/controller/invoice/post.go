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
// @Summary New Invoice
// @Description Creates an invoice using the values supplied on the request body.
// @Tags invoices
// @Accept json
// @Produce  plain
// @Security JwtAuth
// @Param invoice body service.PostModel true "Post Invoice Model. All fields are required."
// @Success 201 {integer} integer "The new invoice Id"
// @Failure 400 {string} string "Indicates a failure when parsing request body or a validation error, e.g. a required field is missing"
// @Failure 401 {string} string "Indicates that no authorization info was provided, or authorization is invalid."
// @Failure 500 {string} string "Indicates an error that was not handled by the server"
// @Router /invoice [post]
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Couldn't create Invoice")
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%v%v%v/%v", r.URL.Scheme, r.Host, r.RequestURI, id))
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, strconv.Itoa(id))
}
