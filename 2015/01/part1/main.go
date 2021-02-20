package main

import (
	"fmt"
	"os"
	"strings"
)

// GetNewFloor Calculates how much floors up and down, and get the new floor
// counting from starting floor
func GetNewFloor(instructions string, startingFloor int) int {
	floorsUp := strings.Count(instructions, "(")
	floorsDown := strings.Count(instructions, ")")
	return startingFloor + floorsUp - floorsDown
}

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	newFloor := GetNewFloor(string(input), 0)
	fmt.Println(newFloor) // GOT 280
}
