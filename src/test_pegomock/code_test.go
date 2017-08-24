package test_pegomock

import (
	"testing"
	. "github.com/petergtz/pegomock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) Test1() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	// This is type checked. It will NOT compile
	// mockPerson.IsOlderThan("dksdsas")
	mockPerson.IsOlderThan(20)

	// Verify
	mockPerson.VerifyWasCalledOnce().IsOlderThan(20)
}

func (suite *TestSuite) Test2() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	mockPerson.IsOlderThan(20)
	mockPerson.IsOlderThan(20)
	mockPerson.IsOlderThan(20)

	// Verify
	mockPerson.VerifyWasCalled(Times(3)).IsOlderThan(20)
}

func (suite *TestSuite) Test3() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	// When stubbing, because of the DSL, the types of the arguments are checked. In gomock/testify mocks they are not.
	// Return values are not type checked but the author of the library agreed to accept PRs from us to add this feature.
	// See commented tests for what happens when the return values are invalid.
	When(mockPerson.IsOlderThan(20)).ThenReturn(true)

	// Verify
	suite.True(mockPerson.IsOlderThan(20))
}

func (suite *TestSuite) TestAdvancedStubbing() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	// In pegomock you can't return values depending on the inputs in these stub callbacks.
	// You can have more complex stubbing in pegomock.
	When(mockPerson.IsOlderThan(20)).Then(func(params []Param) ReturnValues {
		// Do more complicated stuff here.
		// You can test async functions like this for example by adding a clock.Sleep() here
		return []ReturnValue{true}
	})

	suite.True(mockPerson.IsOlderThan(20))
}

// gomock doesn't have arguments capture. This is very useful when your function is called with dynamic input or multiple times
// You can do some table testing and test a slice of inputs against a slice of outputs.
// This might also be useful if your method for example is called with some complex dynamic struct and you want to check it in more detail.
func (suite *TestSuite) TestCapturedArguments() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	mockPerson.IsOlderThan(10)
	mockPerson.IsOlderThan(20)
	mockPerson.IsOlderThan(30)

	// Capture last call only.
	lastSetOfArgs := mockPerson.VerifyWasCalled(AtLeast(1)).IsOlderThan(AnyInt()).GetCapturedArguments()
	suite.Equal(30, lastSetOfArgs)
	// Capture all the calls.
	args := mockPerson.VerifyWasCalled(AtLeast(1)).IsOlderThan(AnyInt()).GetAllCapturedArguments()

	// Arguments are captured as a slice.
	suite.Equal(args, []int{10, 20, 30})
}

// gomock has this feature as well. Only the DSL is different.
func (suite *TestSuite) TestOrderOfCalls() {
	RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	// Please notice they return default values by default.
	mockPerson.GetName()
	mockPerson.GetName()
	mockPerson.IsOlderThan(10)

	inOrderContext := new(InOrderContext)
	mockPerson.VerifyWasCalledInOrder(Twice(), inOrderContext).GetName()
	mockPerson.VerifyWasCalledInOrder(Once(), inOrderContext).IsOlderThan(10)
}

// The Panic message are more helpful than gomock (or testify mocks) when you do the same thing there.
// In gomock the line with the problem is plain text, here it is clickable so you can jump to the line.
//func (suite *TestSuite) TestWhichFailsBecauseOfMoreReturnValuesThanNeeded() {
//	RegisterMockTestingT(suite.T())
//	mockPerson := NewMockPerson()
//
//	When(mockPerson.IsOlderThan(20)).ThenReturn(false, true)
//}
//
//func (suite *TestSuite) TestWhichFailsBecauseOfInvalidReturnValue() {
//	RegisterMockTestingT(suite.T())
//	mockPerson := NewMockPerson()
//
//	When(mockPerson.IsOlderThan(20)).ThenReturn("")
//}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
