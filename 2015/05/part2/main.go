package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsNiceString check if follow the 2 new rules.
func IsNiceString(text string) bool {
	splittedText := strings.Split(text, "")
	atLeastTwice := false
	// Loop the first two letter from the array to the end, if it encounter one,
	// change it to true
	for i := 0; i <= len(splittedText); i++ {
		twoLetters := splittedText[i] + splittedText[i+1]
		if i == len(splittedText)-2 {
			break
		}
		for j := i + 2; j <= len(splittedText)-2; j++ {
			fmt.Println(twoLetters, splittedText[j], splittedText[j+1])
			if twoLetters == splittedText[j]+splittedText[j+1] {
				atLeastTwice = true
				break
			}
		}

		if atLeastTwice == true {
			break
		}
	}

	if !atLeastTwice {
		return false
	}

	// It contains at least one letter which repeats with exactly one letter between them.

	atLeastOneLetter := false
	// AAAA
	for i, val := range splittedText {
		if i+2 >= len(splittedText) {
			break
		}
		if val == splittedText[i+2] {
			atLeastOneLetter = true
		}
	}
	if !atLeastOneLetter {
		return false
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
	fmt.Println(numNiceWords) // GOT 53
}
