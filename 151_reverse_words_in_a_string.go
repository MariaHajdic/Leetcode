package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	for left, right := 0, len(words)-1; left < right; left++ {
		words[left], words[right] = words[right], words[left]
		right--
	}

	return strings.Join(words, " ")
}
