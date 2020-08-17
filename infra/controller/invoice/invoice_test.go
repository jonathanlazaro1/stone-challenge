package invoice

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jonathanlazaro1/stone-challenge/config"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
	"github.com/jonathanlazaro1/stone-challenge/usecase"
)

const endpoint = "/api/v1/invoice"

func makeParameterizedURL(final string) string {
	return fmt.Sprintf("%v/%v", endpoint, final)
}

func TestMain(m *testing.M) {
	config.Load()

	os.Exit(m.Run())
}

func buildRequestFunction(funcBuilder func(service.Invoice) func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	svc := service.NewInvoiceService(usecase.MockInvoiceRepository(100))
	return funcBuilder(*svc)
}

func makeInvoicePostModel() service.PostModel {
	description := "The old description"
	refYear := 2020
	refMonth := 8
	doc := "current doc"
	amount := 2.42

	model := service.PostModel{}
	model.ReferenceYear = &refYear
	model.ReferenceMonth = &refMonth
	model.Document = &doc
	model.Description = &description
	model.Amount = &amount

	return model
}
