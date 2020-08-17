package domain

import (
	"testing"
)

func TestNewInvoice(t *testing.T) {
	sut := NewInvoice()
	isActiveWanted := true

	// A new Invoice should be active
	if isActiveWanted != sut.IsActive {
		t.Errorf("NewInvoice failed, IsActive: got %v, wanted %v", isActiveWanted, sut.IsActive)
	}

	// A new Invoice's CreatedAt should not be zero
	if sut.CreatedAt.IsZero() {
		t.Errorf("NewInvoice failed, CreatedAt: wanted current date and time, got %v", sut.CreatedAt)
	}
}

func TestDeactivatedInvoice(t *testing.T) {
	sut := NewInvoice()
	sut.Deactivate()
	isActiveWanted := false

	// A deactivated Invoice should be... deactivated
	if isActiveWanted != sut.IsActive {
		t.Errorf("DeactivateInvoice failed, IsActive: got %v, wanted %v", isActiveWanted, sut.IsActive)
	}

	// A new Invoice's CreatedAt should not be zero
	if sut.DeactivatedAt.Time.IsZero() {
		t.Errorf("DeactivateInvoice failed, DeactivatedAt: wanted current date and time, got %v", sut.DeactivatedAt.Time)
	}
}
