package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagram_groups := map[string][]string{}

	for _, word := range strs {
		chars := []rune(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		}) // "eat" -> [a,e,t]
		sorted := string(chars)
		anagram_groups[sorted] = append(anagram_groups[sorted], word) // adding the original word
	}

	result := [][]string{}
	for _, group := range anagram_groups {
		result = append(result, group)
	}

	return result
}

/* An overkill approach */ /*
func groupAnagrams(strs []string) [][]string {

}

func isAnagram(word string, letter_bank map[byte]int) bool {
    for _, letter := range word { // checking new word can be created with bank
        if _, ok := letter_bank[letter]; !ok {
            return false // new word has some extra letters
        }
        letter_bank[letter]--
    }

    for _, letter := range letter_bank { // does the map has letters left?
        if letter_bank[letter] != 0 {
            return false
        }
    }

    return true
}*/
