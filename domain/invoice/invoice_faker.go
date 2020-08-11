package invoice

import (
	"time"

	"github.com/brianvoe/gofakeit/v5"
)

func substr(str string, maxLen int) string {
	runes := []rune(str)
	if len(runes) < maxLen {
		return string(runes)
	}
	return string(runes[:maxLen-1])
}

// MakeFakeInvoice generates an Invoice with fake populated data
func MakeFakeInvoice() Invoice {
	in := NewInvoice()
	in.ID = gofakeit.Number(1, 100000)

	currYear := time.Now().Year()
	in.ReferenceYear = gofakeit.Number(currYear-5, currYear-1)
	in.ReferenceMonth = gofakeit.Number(1, 12)

	in.Document = substr(gofakeit.Phrase(), 14)
	in.Description = substr(gofakeit.Phrase(), 256)
	in.Amount = gofakeit.Price(0.01, 1000000)

	return in
}
