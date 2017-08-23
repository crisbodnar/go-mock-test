package test_pegomock

import (
	"testing"
	pm "github.com/petergtz/pegomock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	alien Alien
}

func (suite *TestSuite) SetupTest() {
	suite.alien = Alien{}
}

func (suite *TestSuite) TestUsingPegomock() {
	pm.RegisterMockTestingT(suite.T())
	mockPerson := NewMockPerson()

	// This is type checked. It will NOT compile
	// mockPerson.IsOlderThan("dksdsas")

	suite.alien.isPersonOlderThan(mockPerson, 20)

	// Verify
	mockPerson.VerifyWasCalledOnce().IsOlderThan(20)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
