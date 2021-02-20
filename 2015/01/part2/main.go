package main

import (
	"fmt"
	"os"
)

// GetPositionoOfInstructionThatMakeHimIntoLessOneFloor .That's it.
func GetPositionoOfInstructionThatMakeHimIntoLessOneFloor(instructions string, startingFloor int) int {
	floorsUp := '('
	floorsDown := ')'
	currentPosition := startingFloor
	var finalPosition int
	for pos, instruction := range instructions {
		switch instruction {
		case floorsUp:
			currentPosition++
		case floorsDown:
			currentPosition--
		}
		if currentPosition == -1 {
			finalPosition = pos + 1
			break
		}
	}
	return finalPosition
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	newFloor := GetPositionoOfInstructionThatMakeHimIntoLessOneFloor(string(input), 0)
	fmt.Println(newFloor) // GOT 1797
}
