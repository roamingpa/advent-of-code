package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateMatrix(t *testing.T) {
	cases := make(map[string][][]bool)

	cases[">"] = [][]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}
	cases["^>v<"] = [][]bool{
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
	}

	for input, expected := range cases {
		result := CreateMatrix(input)

		if !cmp.Equal(expected, result) {
			t.Fatalf(`CreateMatrix("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestMarkVisitedHouses(t *testing.T) {
	cases := make(map[string][][]bool)
	cases[">"] = [][]bool{
		{false, false, false},
		{false, true, true},
		{false, false, false},
	}

	cases["^>v<"] = [][]bool{
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, true, true, false, false},
		{false, false, false, true, true, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false},
	}

	for input, expected := range cases {
		matrix := CreateMatrix(input)
		result := MarkVisitedHouses(matrix, input)
		if !cmp.Equal(expected, result) {
			t.Fatalf(`MarkVisitedHouses("%v") = %v, should be %v`, input, result, expected)
		}
	}
}

func TestCountHousesThatReceivesAtLeastOnePresent(t *testing.T) {
	cases := make(map[string]int)
	cases[">"] = 2
	cases["^>v<"] = 4
	cases["^v^v^v^v^v"] = 2

	for input, expected := range cases {
		matrix := CreateMatrix(input)
		visitedHouses := MarkVisitedHouses(matrix, input)
		result := CountHousesThatReceivesAtLeastOnePresent(visitedHouses)
		if !cmp.Equal(expected, result) {
			t.Fatalf(`CountHousesThatReceivesAtLeastOnePresent("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
