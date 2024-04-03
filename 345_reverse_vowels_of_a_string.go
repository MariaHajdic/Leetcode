package main

func reverseVowels(s string) string {
	letters := []byte(s)
	vowels := make(map[byte]bool, 10)
	for _, l := range "aAeEiIoOuU" {
		vowels[byte(l)] = true
	}
	// []byte{65, 69, 73, 79, 85, 97, 101, 105, 111, 117}
	for start, end := 0, len(letters)-1; start < end; {
		if !vowels[letters[start]] {
			start++
			continue
		}
		if !vowels[letters[end]] {
			end--
			continue
		}
		letters[start], letters[end] = letters[end], letters[start]
		start++
		end--
	}
	return string(letters)
}
