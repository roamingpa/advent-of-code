package main

import (
	"testing"
)

func TestParseCases(t *testing.T) {
	parseCases := make(map[string]Dimensions)

	parseCases["2x3x4"] = Dimensions{2, 3, 4}
	parseCases["1x1x10"] = Dimensions{1, 1, 10}

	for input, expected := range parseCases {
		result := ParseDimensions(input)

		if expected != result {
			t.Fatalf(`ParseDimensions("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestGetTotalPaperNeeded(t *testing.T) {
	cases := make(map[string]int)

	cases["2x3x4"] = 58
	cases["1x1x10"] = 43

	for input, expected := range cases {
		parsedDimensions := ParseDimensions(input)
		result := GetTotalPaperNeeded(parsedDimensions)
		if expected != result {
			t.Fatalf(`GetTotalPaperNeeded("%v") = %v, should be %v`, parsedDimensions, result, expected)
		}
	}
}
