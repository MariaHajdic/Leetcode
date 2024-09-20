package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := map[rune]int{}
	for _, letter := range s {
		letters[letter]++
	}
	for _, letter := range t {
		if _, ok := letters[letter]; !ok || letters[letter] < 1 {
			return false
		}
		letters[letter]--
	}
	return true
}
