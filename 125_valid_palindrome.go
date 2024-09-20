package main

import (
	"strings"
	"unicode"
)

func isPalindrome_(s string) bool {
	str := strings.Join(strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}), "")

	for start, end := 0, len(str)-1; start <= end; start++ {
		if str[start] != str[end] {
			return false
		}
		end--
	}

	return true
}
