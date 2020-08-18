package helpers

import (
	"testing"
	"time"
)

type structModel struct {
	Name      *string
	Age       *int
	Active    *bool
	BirthDate *time.Time
}

func TestCopyIfNotNil(t *testing.T) {
	src := structModel{
		Active:    new(bool),
		BirthDate: new(time.Time),
	}
	dest := structModel{
		Name: new(string),
		Age:  new(int),
	}

	CopyIfNotNil(&src, &dest)

	if dest.Active == nil || dest.BirthDate == nil {
		t.Fatalf("CopyIfNotNil failed: wanted struct fully copied, got %v", dest)
	}

}
