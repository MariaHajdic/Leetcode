package main

func canConstruct(ransomNote string, magazine string) bool {
	alphabet := make(map[rune]int, 26)
	for _, letter := range magazine {
		alphabet[letter] += 1
	}
	for _, letter := range ransomNote {
		alphabet[letter] -= 1
		if alphabet[letter] < 0 {
			return false
		}
	}
	return true
}
