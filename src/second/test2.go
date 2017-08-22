package second

import (
	"testing"
	pm "github.com/petergtz/pegomock"
)

func TestUsingPegomock(t *testing.T) {
	pm.RegisterMockTestingT(t)
	mockPerson := NewMockPerson()

	// This is type checked. It will NOT compile
	// mockPerson.IsOlderThan("dksdsas")

	// Using the mock
	mockPerson.IsOlderThan(20)

	// Verify
	mockPerson.VerifyWasCalledOnce().IsOlderThan(20)
}