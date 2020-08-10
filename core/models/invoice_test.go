package models

import (
	"testing"
)

func TestNewInvoice(t *testing.T) {
	sut := NewInvoice()
	isActiveWanted := true

	// A new Invoice should be active
	if isActiveWanted != sut.IsActive {
		t.Errorf("NewInvoice failed: IsActive should be %v, but was %v", isActiveWanted, sut.IsActive)
	}

	// A new Invoice's CreatedAt should not be zero
	if sut.CreatedAt.IsZero() {
		t.Errorf("NewInvoice failed: CreatedAt should be the current date and time, but was %v", sut.CreatedAt)
	}
}
