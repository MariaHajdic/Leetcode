package main

// func strStr(haystack string, needle string) int {
//     return strings.Index(haystack, needle)
// }

/* More fun solution. */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
