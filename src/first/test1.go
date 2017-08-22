package first

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestUsingPegomock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	// Mock expectations are asserted on Finish() call of the controller.
	defer mockCtrl.Finish()

	mockPerson := NewMockPerson(mockCtrl)
	// This is not typed checked and it would compile. Only the number of arguments is checked.
	// mockPerson.EXPECT().IsOlderThan("Twenty")

	// This is the correct version.
	mockPerson.EXPECT().IsOlderThan(20)

	// Return values are not typed checked. This would compile.
	// mockPerson.EXPECT().IsOlderThan(30).Return("invalid return", 10202, []string{})

	// This is a correct version.
	mockPerson.EXPECT().IsOlderThan(20).Return(true).Times(2)

	// Check the order methods are called
	gomock.InOrder(
		mockPerson.EXPECT().IsOlderThan(20),
		mockPerson.EXPECT().GetName(),
	)

	// Chain expected calls.
	firstCall := mockPerson.EXPECT().IsOlderThan(10)
	secondCall := mockPerson.EXPECT().IsOlderThan(20).After(firstCall)
	_ := mockPerson.EXPECT().IsOlderThan(30).After(secondCall)

	// Argument matchers. Custom argument matchers are supported as well
	mockPerson.EXPECT().IsOlderThan(gomock.Eq(19)).Return(true).Times(2)

	// Actions for more complex stubs.
	mockPerson.EXPECT().
		IsOlderThan(gomock.Any()).
		Return(nil).
		// Do function is not type checked.
		Do(func(age int) {
		if age > 10 {
			t.Fail()
		}
	})

	// Using the mock
	mockPerson.GetName()
	mockPerson.IsOlderThan(20)
}
