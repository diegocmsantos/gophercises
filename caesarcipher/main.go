package main

import (
	"fmt"
	"strings"
	"unicode"
)

func caesar(r rune, shift int32) rune {

	if !unicode.IsLetter(r) {
		return r
	}

	s := int32(r) + shift
	if unicode.IsLower(r) {
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
	}
	return rune(s)
}

func caesarcipher(s string, k int32) string {
	return strings.Map(func(r rune) rune {
		return caesar(r, k)
	}, s)
}

func main() {
	fmt.Println(caesarcipher("middle-Outz", 2))
}
