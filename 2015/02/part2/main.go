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

// ParseDimensions of the box into the struct Dimensions
func ParseDimensions(dimensions string) Dimensions {
	splittedDimensions := strings.Split(dimensions, "x")
	length, _ := strconv.Atoi(splittedDimensions[0])
	width, _ := strconv.Atoi(splittedDimensions[1])
	height, _ := strconv.Atoi(splittedDimensions[2])

	parsedDimensions := Dimensions{length, width, height}
	return parsedDimensions
}

// GetVolumeOfBox for bow size.
func GetVolumeOfBox(dimensions Dimensions) int {
	volume := dimensions.height * dimensions.length * dimensions.width
	return volume
}

// GetSmallestPerimeterOfABox for ribbon size.
func GetSmallestPerimeterOfABox(dimensions Dimensions) int {
	sides := []int{dimensions.height, dimensions.length, dimensions.width}
	var maxDimensionIndex int
	for i, value := range sides {
		if sides[i] == 0 {
			maxDimensionIndex = i
		}
		if value > sides[maxDimensionIndex] {
			maxDimensionIndex = i
		}
	}
	smallestSides := append(sides[:maxDimensionIndex], sides[maxDimensionIndex+1:]...)

	smallestPerimeter := (smallestSides[0] * 2) + (smallestSides[1] * 2)
	return smallestPerimeter

}

// GetTotalRibbonNeeded for a box.
func GetTotalRibbonNeeded(dimensions Dimensions) int {
	return GetVolumeOfBox(dimensions) + GetSmallestPerimeterOfABox(dimensions)
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
		totalRibbon := GetTotalRibbonNeeded(dimensions)
		needToOrder += totalRibbon
	}
	fmt.Println(needToOrder) // GOT 3842356
}
