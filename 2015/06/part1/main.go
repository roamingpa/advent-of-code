package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Position ...
type Position struct {
	x int
	y int
}

var lights = make(map[Position]bool)

// GetPointsFromInstruction ...
func GetPointsFromInstruction(instruction string) ([]Position, string) {
	stringToReplace := ""
	if strings.Contains(instruction, "turn on ") {
		stringToReplace = "turn on "
	}
	if strings.Contains(instruction, "turn off ") {

		stringToReplace = "turn off "
	}
	if strings.Contains(instruction, "toggle ") {
		stringToReplace = "toggle "

	}

	instruction = strings.Replace(instruction, stringToReplace, "", 1)
	firstCommaIndex := strings.Index(instruction, ",")
	xValue := instruction[0:firstCommaIndex]
	x, err := strconv.Atoi(xValue)
	if err != nil {
		panic(err)
	}
	instruction = instruction[firstCommaIndex+1:]
	nextSpace := strings.Index(instruction, " ")
	yValue := instruction[0:nextSpace]
	instruction = instruction[nextSpace+1:]
	y, err := strconv.Atoi(yValue)
	if err != nil {
		panic(err)
	}
	initialPosition := Position{x: x, y: y}

	// Final Position
	instruction = strings.Replace(instruction, "through ", "", 1)
	secondCommaIndex := strings.Index(instruction, ",")
	xValue = instruction[0:secondCommaIndex]
	x, err = strconv.Atoi(xValue)
	if err != nil {
		panic(err)
	}
	instruction = instruction[secondCommaIndex+1:]
	yValue = instruction[0:]
	y, err = strconv.Atoi(yValue)
	if err != nil {
		panic(err)
	}
	finalPosition := Position{x: x, y: y}

	return []Position{initialPosition, finalPosition}, stringToReplace
}

func setUpLights(positions []Position, action string) {
	if strings.Contains(action, "turn on ") {
		for i := positions[0].x; i <= positions[1].x; i++ {
			for j := positions[0].y; j <= positions[1].y; j++ {
				lights[Position{i, j}] = true
			}
		}
	}
	if strings.Contains(action, "toggle ") {
		for i := positions[0].x; i <= positions[1].x; i++ {
			for j := positions[0].y; j <= positions[1].y; j++ {
				lights[Position{i, j}] = !lights[Position{i, j}]
			}
		}
	}
	if strings.Contains(action, "turn off ") {
		// Initial Position
		for i := positions[0].x; i <= positions[1].x; i++ {
			for j := positions[0].y; j <= positions[1].y; j++ {
				lights[Position{i, j}] = false
			}
		}
	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		points, action := GetPointsFromInstruction(scanner.Text())
		setUpLights(points, action)
		// fmt.Println(lights)
	}
	lightsLit := 0
	for _, val := range lights {
		if val == true {
			lightsLit++
		}
	}
	fmt.Println(lightsLit) // GOT 400410
}
