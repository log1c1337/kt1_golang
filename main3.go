package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	reversed := make([]rune, n)

	for i, r := range runes {
		reversed[n-1-i] = r
	}

	return string(reversed)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевёрнутая строка:", reversed)
}
