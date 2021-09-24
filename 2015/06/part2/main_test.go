package main

import (
	"testing"
)

func TestGetPointsFromInstruction(t *testing.T) {
	cases := make(map[string][]Position)

	cases["turn off 660,55 through 986,197"] = []Position{{660, 55}, {986, 197}}
	cases["turn on 226,196 through 599,390"] = []Position{{226, 196}, {599, 390}}
	cases["toggle 767,98 through 854,853"] = []Position{{767, 98}, {854, 853}}
	for input, expected := range cases {
		result, _ := GetPointsFromInstruction(input)

		if result[0].x != expected[0].x || result[0].y != expected[0].y || result[1].x != expected[1].x || result[1].y != expected[1].y {
			t.Fatalf(`getPointsFromInstruction("%v") = %v, should be %v`, input, result, expected)
		}
	}
}
