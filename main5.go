package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {

	s = strings.ToLower(s)

	charMap := make(map[rune]bool)

	for _, char := range s {

		if charMap[char] {
			return false
		}

		charMap[char] = true
	}
	return true
}

func main() {
	fmt.Println(hasUniqueChars("abcd"))
	fmt.Println(hasUniqueChars("abCdefAaf"))
	fmt.Println(hasUniqueChars("aabcd"))
}
