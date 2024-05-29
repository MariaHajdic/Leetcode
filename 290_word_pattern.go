package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	letterToWord := map[byte]string{}
	wordToLetter := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		letter := pattern[i]
		word := words[i]

		if mappedWord, ok := letterToWord[letter]; ok {
			if mappedWord != word {
				return false
			}
			continue
		}
		if mappedLetter, ok := wordToLetter[word]; ok {
			if mappedLetter != letter {
				return false // bijections only!
			}
		}
		letterToWord[letter] = word
		wordToLetter[word] = letter
	}

	return true
}
