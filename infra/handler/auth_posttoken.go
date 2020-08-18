package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// AuthPostTokenHandler processes requests to authentication
func AuthPostTokenHandler(w http.ResponseWriter, r *http.Request) {
	var model AuthRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errCouldntParseAuthModel)
		return
	}

	err = model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint(err))
		return
	}

	token, err := service.Authenticate(model.Email, model.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		io.WriteString(w, errCouldntGenerateToken)
		return
	}

	json.NewEncoder(w).Encode(&AuthResponseModel{Token: token})

}
