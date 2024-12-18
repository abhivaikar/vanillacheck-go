package main

import (
	"vanillacheck-go/vanillacheck"

	"github.com/stretchr/testify/assert"
)

type TestSuite struct{}

func (ts *TestSuite) BeforeTest(tc *vanillacheck.TestContext) {
	// Setup logic before each test
}

func (ts *TestSuite) AfterTest(tc *vanillacheck.TestContext) {
	// Cleanup logic after each test
}

func (ts *TestSuite) TestAddition(tc *vanillacheck.TestContext) {
	assert.Equal(tc, 4, 2+2, "2 + 2 should equal 4")
}

func (ts *TestSuite) TestFailure(tc *vanillacheck.TestContext) {
	tc.Errorf("Intentional failure")
}
