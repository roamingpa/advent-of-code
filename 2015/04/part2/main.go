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

// FindFirstHashWith6LeadingZeros just put 1 more zero...
func FindFirstHashWith6LeadingZeros(input string) string {
	i := 0
	for {
		hashed := CreateMD5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(hashed, "000000") {
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
	result := FindFirstHashWith6LeadingZeros(string(input))
	fmt.Println(result) // GOT 3938038
}
