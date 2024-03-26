package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	// Iterating through words.
	for i := 1; i < len(strs) && len(prefix) > 0; i++ {
		word := strs[i]

		// Iterating through letters in current word.
		j := 0
		for ; j < len(word) && j < len(prefix); j++ {
			if prefix[j] != word[j] {
				break
			}
		}
		prefix = prefix[:j]
	}

	return prefix
}
