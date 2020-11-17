package utils

import (
	"fmt"
	"testing"
)

var trimtests = map[string]string{
	"///string///":                 "///string",
	"http://www.example.com/test/": "http://www.example.com/test",
	"string":                       "string",
}

var emptytests = map[string]bool{
	"":       true,
	"abc123": false,
	"🤘🏻":     false,
}

func TestStringTrimming(t *testing.T) {
	for original, expected := range trimtests {
		if result := Sanatize(original); result != expected {
			t.Logf("Expected %q but received %q instead.", expected, result)
			t.Fail()
		}
	}
}

func TestEmpties(t *testing.T) {
	for value, expected := range emptytests {
		if empty := IsEmpty(value); empty != expected {
			t.Logf(
				"String %q expected empty to be %q, got %q instead",
				value, fmt.Sprintf("%v", expected), fmt.Sprintf("%v", empty),
			)
			t.Fail()
		}
	}
}
