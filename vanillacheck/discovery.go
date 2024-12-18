package vanillacheck

import (
	"reflect"
	"strings"
)

// DiscoverTests identifies all methods starting with Test* in the given suite.
func DiscoverTests(testSuite interface{}) []string {
	var tests []string
	typ := reflect.TypeOf(testSuite)

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		if strings.HasPrefix(method.Name, "Test") && method.Type.NumIn() == 2 {
			tests = append(tests, method.Name)
		}
	}
	return tests
}
