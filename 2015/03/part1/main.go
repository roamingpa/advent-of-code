package main

import (
	"fmt"
	"os"
	"strings"
)

// RemoveCircularMoves aaa
func RemoveCircularMoves(instructions string) string {
	removedUpDown := strings.ReplaceAll(instructions, "^v", "")
	removedDownUp := strings.ReplaceAll(removedUpDown, "v^", "")
	removedRightLeft := strings.ReplaceAll(removedDownUp, "><", "")
	removedLeftRight := strings.ReplaceAll(removedRightLeft, "<>", "")

	return removedLeftRight
}

// CreateMatrix by getting the minLegth needed.
func CreateMatrix(instructions string) [][]bool {
	minLength := len(instructions) + 2
	if minLength%2 == 0 {
		minLength++
	}
	matrix := make([][]bool, minLength)

	for i := 0; i < minLength; i++ {
		matrix[i] = make([]bool, minLength)
		for j := 0; j < minLength; j++ {
			if i == (minLength/2) && j == (minLength/2) {
				matrix[i][j] = true
			} else {
				matrix[i][j] = false
			}
		}
	}
	return matrix
}

// MarkVisitedHouses by changing the state.
func MarkVisitedHouses(matrix [][]bool, instructions string) [][]bool {
	currentX := len(matrix) / 2
	currentY := len(matrix) / 2

	for _, val := range strings.Split(instructions, "") {
		switch val {
		case "^":
			currentY--
		case ">":
			currentX++
		case "v":
			currentY++
		case "<":
			currentX--
		}
		matrix[currentY][currentX] = true
	}
	return matrix
}

// CountHousesThatReceivesAtLeastOnePresent checking the matrix bools.
func CountHousesThatReceivesAtLeastOnePresent(matrix [][]bool) int {
	numHouses := 0
	for _, row := range matrix {
		for _, val := range row {
			if val == true {
				numHouses++
			}
		}
	}
	return numHouses
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	matrix := CreateMatrix(string(input))
	visitedHouses := MarkVisitedHouses(matrix, string(input))
	numHouses := CountHousesThatReceivesAtLeastOnePresent(visitedHouses)
	fmt.Println(numHouses) // Got 2081
}
