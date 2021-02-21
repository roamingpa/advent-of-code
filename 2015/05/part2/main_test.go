package main

import (
	"testing"
)

func TestIsNiceString(t *testing.T) {
	cases := make(map[string]bool)
	cases["qjhvhtzxzqqjkmpb"] = true
	cases["xxyxx"] = true
	cases["uurcxstgmygtbstg"] = false
	cases["ieodomkazucvgmuy"] = false

	for input, expected := range cases {
		result := IsNiceString(input)

		if result != expected {
			t.Fatalf(`IsNiceString("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
