package main

func lengthOfLastWord(s string) int {
	letters := []byte(s) // for ascii only
	length := 0
	for i := len(letters) - 1; i >= 0; i-- {
		if letters[i] != ' ' {
			length++
			continue
		}
		if length > 0 {
			break
		}
	}
	return length
}

/* Doesn't cover multiple spaces.
func lengthOfLastWord(s string) int {
    words := strings.Split(s, " ")
    return len(words[len(words)-1])
}
*/
