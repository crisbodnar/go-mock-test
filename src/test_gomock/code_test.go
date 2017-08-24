package test_gomock

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

func (t *TestSuite) TestTheBasics() {
	mockCtrl := gomock.NewController(t.T())
	// Mock expectations are asserted on Finish() call of the controller.
	defer mockCtrl.Finish()

	mockPerson := NewMockPerson(mockCtrl)
	// This is not typed checked and it would compile. Only the number of arguments is checked.
	// mockPerson.EXPECT().IsOlderThan("Twenty")

	// This is a working version.
	mockPerson.EXPECT().IsOlderThan(20)

	// This is type checked.
	mockPerson.IsOlderThan(20)
}

//func TestWithInvalidReturnValues(t *testing.T) {
//	mockCtrl := gomock.NewController(t)
//	defer mockCtrl.Finish()
//
//	mockPerson := NewMockPerson(mockCtrl)
//
//	//// Return values are not typed checked. This compiles.
//	mockPerson.EXPECT().IsOlderThan(30).Return("invalid return", 10202, []string{})
//	mockPerson.IsOlderThan(30)
//}

func (t *TestSuite) TestStubbing() {
	mockCtrl := gomock.NewController(t.T())
	defer mockCtrl.Finish()

	mockPerson := NewMockPerson(mockCtrl)

	//// Return values are not typed checked. This compiles.
	mockPerson.EXPECT().IsOlderThan(30).Return(true)
	t.True(mockPerson.IsOlderThan(30))
}

// A big difference from pegomock. Methods are not stubbed by default. This will panic. In pegomock it would return default values
//func (t *TestSuite) TestStubbingWhenNoExpectationDefined() {
//	mockCtrl := gomock.NewController(t.T())
//	defer mockCtrl.Finish()
//
//	mockPerson := NewMockPerson(mockCtrl)
//
//	t.True(mockPerson.IsOlderThan(30))
//}

func (t *TestSuite) TestOrderOfCalls() {
	mockCtrl := gomock.NewController(t.T())
	defer mockCtrl.Finish()

	mockPerson := NewMockPerson(mockCtrl)

	gomock.InOrder(
		mockPerson.EXPECT().IsOlderThan(20).Times(2),
		mockPerson.EXPECT().GetName(),
	)

	mockPerson.IsOlderThan(20)
	mockPerson.IsOlderThan(20)
	mockPerson.GetName()
}

func (t *TestSuite) TestStubsCallbacks() {
	mockCtrl := gomock.NewController(t.T())
	defer mockCtrl.Finish()

	mockPerson := NewMockPerson(mockCtrl)

	mockPerson.EXPECT().
		IsOlderThan(gomock.Any()).Times(2).
		Return(true).
		Do(func(age int) {
		// More advanced stubbing. Problem is you can't return values depending on the input as in pegomock.
		// You can at most validate more complex inputs and panic if something is not right.
		if age > 10 {
			println(age)
		}
	})

	t.True(mockPerson.IsOlderThan(53))
	t.True(mockPerson.IsOlderThan(45))
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

