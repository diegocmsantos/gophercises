package main

import (
	"fmt"
	"strings"
	"unicode"
)

func camelcase(s string) int32 {
	isUpperFn := func(r rune) bool {
		return unicode.IsUpper(r)
	}
	words := strings.FieldsFunc(s, isUpperFn)

	return int32(len(words))

}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
