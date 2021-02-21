package main

import (
	"testing"
)

func TestCreateMD5Hash(t *testing.T) {
	cases := make(map[string]string)
	cases["abcdef609043"] = "000001dbbfa3a5c83a2d506429c7b00e"
	cases["pqrstuv1048970"] = "000006136ef2ff3b291c85725f17325c"

	for input, expected := range cases {
		result := CreateMD5Hash(input)

		if result != expected {
			t.Fatalf(`CreateMD5Hash("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestFindFirstHashWith5LeadingZeros(t *testing.T) {
	cases := make(map[string]string)
	cases["abcdef"] = "609043"
	cases["pqrstuv"] = "1048970"

	for input, expected := range cases {
		result := FindFirstHashWith5LeadingZeros(input)

		if result != expected {
			t.Fatalf(`FindFirstHashWith5LeadingZeros("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
