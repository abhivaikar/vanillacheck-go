package vanillacheck

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"time"
)

type TestResult struct {
	Name       string
	Passed     bool
	ErrorMsg   string
	StackTrace string
	Runtime    time.Duration
}

type TestRunner struct {
	Results []TestResult
}

func NewTestRunner() *TestRunner {
	return &TestRunner{}
}

func (tr *TestRunner) RunTests(testSuite interface{}) {
	testNames := DiscoverTests(testSuite)
	val := reflect.ValueOf(testSuite) // This is used to call methods dynamically.

	for _, testName := range testNames {
		start := time.Now()
		result := TestResult{Name: testName, Passed: true}
		tc := NewTestContext()

		// Run BeforeTest hook if available
		if beforeMethod := val.MethodByName("BeforeTest"); beforeMethod.IsValid() {
			beforeMethod.Call([]reflect.Value{reflect.ValueOf(tc)})
			if tc.HasFailed() {
				result.Passed = false
				result.ErrorMsg = tc.FailureMessage()
				tr.Results = append(tr.Results, result)
				continue
			}
		}

		// Execute the test method
		func() {
			defer func() {
				if r := recover(); r != nil {
					result.Passed = false
					result.ErrorMsg = fmt.Sprintf("panic: %v", r)
					result.StackTrace = string(debug.Stack())
				}
			}()
			method := val.MethodByName(testName)
			method.Call([]reflect.Value{reflect.ValueOf(tc)})
		}()

		// Check test failure
		if tc.HasFailed() {
			result.Passed = false
			result.ErrorMsg = tc.FailureMessage()
		}

		// Run AfterTest hook if available
		if afterMethod := val.MethodByName("AfterTest"); afterMethod.IsValid() {
			afterMethod.Call([]reflect.Value{reflect.ValueOf(tc)})
		}

		result.Runtime = time.Since(start)
		tr.Results = append(tr.Results, result)
	}

	tr.PrintSummary()
	tr.WriteJSONReport("results.json")
	tr.WriteHTMLReport("results.html")
}
