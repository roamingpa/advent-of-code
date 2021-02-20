package main

import (
	"testing"
)

func TestBaseCases(t *testing.T) {
	var cases map[string]int
	cases = make(map[string]int)
	startingFloor := 0
	cases["(())"] = 0
	cases["()()"] = 0

	cases["((("] = 3
	cases["(()(()("] = 3
	cases["))((((("] = 3

	cases["())"] = -1
	cases["))("] = -1

	cases[")))"] = -3
	cases[")())())"] = -2

	for input, expected := range cases {
		result := GetNewFloor(input, startingFloor)

		if expected != result {
			t.Fatalf(`GetNewFloor("%v", %v) = %v, should be %v`, input, startingFloor, result, expected)
		}
	}
}
