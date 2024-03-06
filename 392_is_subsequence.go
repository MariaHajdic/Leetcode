package main

func isSubsequence(s string, t string) bool {
	idx1, idx2 := 0, 0
	for ; idx1 < len(s) && idx2 < len(t); idx2++ {
		if s[idx1] == t[idx2] {
			idx1++
		}
	}

	return idx1 > len(s)-1
}
