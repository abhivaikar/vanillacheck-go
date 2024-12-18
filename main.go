package main

import (
	"vanillacheck-go/vanillacheck"
)

func main() {
	runner := vanillacheck.NewTestRunner()
	runner.RunTests(&TestSuite{})
}
