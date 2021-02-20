package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Dimensions of the box
type Dimensions struct {
	length int
	width  int
	height int
}

// GetTotalPaperNeeded Calc the total paper needed, getting the area of
// the rectangular prism plus the extra smallest side.
func GetTotalPaperNeeded(dimensions Dimensions) int {
	prismArea := 0

	side1 := 2 * dimensions.length * dimensions.width
	side2 := 2 * dimensions.width * dimensions.height
	side3 := 2 * dimensions.height * dimensions.length

	sides := []int{side1, side2, side3}

	var smallestSide int
	for i, val := range sides {
		prismArea += val
		if i == 0 {
			smallestSide = val
		}
		if smallestSide > val {
			smallestSide = val
		}
	}
	totalPaperNeeded := prismArea + smallestSide/2
	return totalPaperNeeded
}

// ParseDimensions of the box into the struct Dimensions
func ParseDimensions(dimensions string) Dimensions {
	splittedDimensions := strings.Split(dimensions, "x")
	length, _ := strconv.Atoi(splittedDimensions[0])
	width, _ := strconv.Atoi(splittedDimensions[1])
	height, _ := strconv.Atoi(splittedDimensions[2])

	parsedDimensions := Dimensions{length, width, height}
	return parsedDimensions
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	needToOrder := 0
	for scanner.Scan() {
		dimensions := ParseDimensions(scanner.Text())
		totalPaper := GetTotalPaperNeeded(dimensions)
		needToOrder += totalPaper
	}
	fmt.Println(needToOrder) // GOT 1606483
}
