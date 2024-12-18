package vanillacheck

import "fmt"

// TestContext provides methods for test assertions and logging.
type TestContext struct {
	failed  bool
	message string
}

// NewTestContext creates a new TestContext instance.
func NewTestContext() *TestContext {
	return &TestContext{}
}

// Fail marks the test as failed but allows execution to continue.
func (tc *TestContext) Fail() {
	tc.failed = true
}

// FailNow marks the test as failed and stops execution immediately.
func (tc *TestContext) FailNow() {
	tc.failed = true
	panic("FailNow called")
}

// Error logs an error message and marks the test as failed.
func (tc *TestContext) Error(args ...interface{}) {
	tc.failed = true
	tc.message = fmt.Sprint(args...)
}

// Errorf logs a formatted error message and marks the test as failed.
func (tc *TestContext) Errorf(format string, args ...interface{}) {
	tc.failed = true
	tc.message = fmt.Sprintf(format, args...)
}

// Fatal logs an error message, marks the test as failed, and stops execution immediately.
func (tc *TestContext) Fatal(args ...interface{}) {
	tc.failed = true
	tc.message = fmt.Sprint(args...)
	panic("Fatal called")
}

// Fatalf logs a formatted error message, marks the test as failed, and stops execution immediately.
func (tc *TestContext) Fatalf(format string, args ...interface{}) {
	tc.failed = true
	tc.message = fmt.Sprintf(format, args...)
	panic("Fatalf called")
}

// HasFailed checks if the test has failed.
func (tc *TestContext) HasFailed() bool {
	return tc.failed
}

// FailureMessage returns the failure message, if any.
func (tc *TestContext) FailureMessage() string {
	return tc.message
}
