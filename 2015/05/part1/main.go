package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsNiceString check if follow the 3 rules.
func IsNiceString(text string) bool {
	// At least 3 vowels
	totalVowels := 0
	totalVowels += strings.Count(text, "a")
	totalVowels += strings.Count(text, "e")
	totalVowels += strings.Count(text, "i")
	totalVowels += strings.Count(text, "o")
	totalVowels += strings.Count(text, "u")
	if totalVowels < 3 {
		return false
	}
	splittedText := strings.Split(text, "")

	// At least one letter that appears twice in a row
	hasOneLetterTwiceInARow := false
	for i, val := range splittedText {
		if i > 0 && splittedText[i-1] == val {
			hasOneLetterTwiceInARow = true
			break
		}
	}
	if !hasOneLetterTwiceInARow {
		return false
	}
	// Doesn't containt the strings ab, cd, pq, xy
	cantContain := []string{"ab", "cd", "pq", "xy"}
	for _, val := range cantContain {
		if strings.Contains(text, val) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numNiceWords := 0
	for scanner.Scan() {
		if IsNiceString(scanner.Text()) {
			numNiceWords++
		}
	}
	fmt.Println(numNiceWords) // GOT 258
}
