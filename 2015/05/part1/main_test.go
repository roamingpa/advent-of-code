package main

import (
	"testing"
)

func TestIsNiceString(t *testing.T) {
	cases := make(map[string]bool)
	cases["ugknbfddgicrmopn"] = true
	cases["aaa"] = true
	cases["jchzalrnumimnmhp"] = false
	cases["haegwjzuvuyypxyu"] = false
	cases["dvszwmarrgswjxmb"] = false

	for input, expected := range cases {
		result := IsNiceString(input)

		if result != expected {
			t.Fatalf(`IsNiceString("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
