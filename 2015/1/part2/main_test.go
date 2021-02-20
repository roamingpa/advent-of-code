package main

import (
	"testing"
)

func TestBaseCases(t *testing.T) {
	var cases map[string]int
	cases = make(map[string]int)
	startingFloor := 0
	cases[")"] = 1
	cases["()())"] = 5

	for input, expected := range cases {
		result := GetPositionoOfInstructionThatMakeHimIntoLessOneFloor(input, startingFloor)

		if expected != result {
			t.Fatalf(`GetPositionoOfInstructionThatMakeHimIntoLessOneFloor("%v", %v) = %v, should be %v`, input, startingFloor, result, expected)
		}
	}
}
