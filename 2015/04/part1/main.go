package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CreateMD5Hash from an input
func CreateMD5Hash(input string) string {
	hashInBytesArray := md5.Sum([]byte(input))
	return hex.EncodeToString(hashInBytesArray[:])
}

// FindFirstHashWith5LeadingZeros ...
func FindFirstHashWith5LeadingZeros(input string) string {
	i := 0
	for {
		hashed := CreateMD5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(hashed, "00000") {
			return strconv.Itoa(i)
		}
		i++
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	result := FindFirstHashWith5LeadingZeros(string(input))
	fmt.Println(result) // GOT 117946
}
