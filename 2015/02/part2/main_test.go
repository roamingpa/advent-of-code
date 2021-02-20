package main

import (
	"testing"
)

func TestGetSmallestPerimeter(t *testing.T) {
	cases := make(map[string]int)

	cases["2x3x4"] = 10
	cases["1x1x10"] = 4

	for input, expected := range cases {
		parsedDimensions := ParseDimensions(input)
		result := GetSmallestPerimeterOfABox(parsedDimensions)

		if expected != result {
			t.Fatalf(`GetSmallestPerimeterOfABox("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestGetVolumeOfBox(t *testing.T) {
	cases := make(map[string]int)

	cases["2x3x4"] = 24
	cases["1x1x10"] = 10

	for input, expected := range cases {
		parsedDimensions := ParseDimensions(input)
		result := GetVolumeOfBox(parsedDimensions)

		if expected != result {
			t.Fatalf(`GetVolumeOfBox("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestGetTotalRibbonNeeded(t *testing.T) {
	cases := make(map[string]int)

	cases["2x3x4"] = 34
	cases["1x1x10"] = 14

	for input, expected := range cases {
		parsedDimensions := ParseDimensions(input)
		result := GetTotalRibbonNeeded(parsedDimensions)

		if expected != result {
			t.Fatalf(`GetTotalRibbonNeeded("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
