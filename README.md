
# vanillacheck-go

Hi there! I'm working on this project, **vanillacheck-go**, as an exercise to better understand how testing frameworks like Go's `testing` package or JUnit are created. This is not meant to compete with those frameworks; it's just a simple and experimental attempt to build something from scratch to learn the core concepts behind test discovery, execution, and reporting.

---

## Why I Created vanillacheck-go

I’ve always been fascinated by how popular testing frameworks make writing and running tests so seamless. To understand the inner workings, I decided to create my own basic test runner. The goal is to explore:
- **How test methods are discovered automatically using reflection.**
- **How setup and teardown hooks are implemented.**
- **How test results are captured and presented.**

This project is my way of diving deeper into those concepts and experimenting with a lightweight test runner in Go.

---

## What vanillacheck-go Does

vanillacheck-go is a minimal test runner that:
1. Automatically discovers test methods in a `TestSuite` struct.
2. Executes tests sequentially with `BeforeTest` and `AfterTest` hooks.
3. Generates test result reports in three formats:
    - Console output
    - JSON file
    - HTML file

---

## Features

- **Automatic Test Discovery**:
    - Test methods are identified by the `Test` prefix in a `TestSuite` struct.

- **Simple Test Execution**:
    - Tests are executed sequentially.
    - Failures in one test don’t stop others from running.

- **Custom Assertions**:
    - A `TestContext` provides assertion methods like `Fail`, `Error`, and `Fatal`.

- **Report Generation**:
    - A JSON file summarizing the results.
    - An HTML report for easy viewing.

---

## How It Works

Here’s a simple example of how to use vanillacheck-go:

### Write a TestSuite
```go
package main

import (
	"vanillacheck-go/vanillacheck"
	"github.com/stretchr/testify/assert"
)

type TestSuite struct{}

func (ts *TestSuite) BeforeTest(tc *vanillacheck.TestContext) {
	// Setup logic
}

func (ts *TestSuite) AfterTest(tc *vanillacheck.TestContext) {
	// Teardown logic
}

func (ts *TestSuite) TestAddition(tc *vanillacheck.TestContext) {
	assert.Equal(tc, 4, 2+2, "2 + 2 should equal 4")
}

func (ts *TestSuite) TestFailure(tc *vanillacheck.TestContext) {
	tc.Errorf("This test is intentionally failing")
}
```

### Run the Tests
The `TestRunner` class is responsible for running your tests:
```go
package main

import (
	"vanillacheck-go/vanillacheck"
)

func main() {
	runner := vanillacheck.NewTestRunner()
	runner.RunTests(&TestSuite{})
}
```

---

## What You’ll See

### **1. Console Output**
vanillacheck-go prints a summary of the test results in the console:
```plaintext
=== Test Summary ===
Total tests run: 2
Passed: 1
Failed: 1
====================

Detailed Results:
Test: TestAddition | Status: PASS
Test: TestFailure  | Status: FAIL
Reason: This test is intentionally failing
```

### **2. JSON Report (`results.json`)**
The results are also written to a JSON file for structured analysis:
```json
[
    {
        "testName": "TestAddition",
        "status": "PASS",
        "errorMessage": null,
        "executionTime": 5
    },
    {
        "testName": "TestFailure",
        "status": "FAIL",
        "errorMessage": "This test is intentionally failing",
        "executionTime": 7
    }
]
```

### **3. HTML Report (`results.html`)**
The HTML report provides a visually appealing summary of test results:
- **Total tests**, **passed**, and **failed** are displayed in a table.
- Each test includes its name, status, error message (if any), and execution time.
