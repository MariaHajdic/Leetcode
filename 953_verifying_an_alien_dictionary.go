package main

func isAlienSorted(words []string, order string) bool {
	letters := map[byte]int{} // letters # in alphabet
	for i := range len(order) {
		letters[order[i]] = i
	}

	prev_word := words[0]
	for i := 1; i < len(words); i++ {
		if !sorted(prev_word, words[i], letters) {
			return false
		}
		prev_word = words[i]
	}
	return true
}

func sorted(word1, word2 string, order map[byte]int) bool {
	min_len := len(word1)
	if len(word2) < min_len {
		min_len = len(word2)
	}

	for i := 0; i < min_len; i++ {
		if order[word1[i]] < order[word2[i]] { // definitely sorted
			return true
		} else if order[word1[i]] > order[word2[i]] { // definitely not
			return false
		}
	}

	return len(word1) <= len(word2)
}
