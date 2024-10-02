package main

func findRepeatedDnaSequences(s string) []string {
	repeated := map[string]int{}
	result := []string{}

	for left, right := 0, 9; right < len(s); right++ {
		if repeated[s[left:right+1]] == 1 { // already found before
			result = append(result, s[left:right+1])
		}
		repeated[s[left:right+1]]++
		left++
	}

	return result
}
