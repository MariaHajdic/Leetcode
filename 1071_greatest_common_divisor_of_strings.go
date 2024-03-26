package main

import "strings"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	greatestDivisor := gcd(len(str1), len(str2))
	pattern := str1[:greatestDivisor]

	if strings.Count(str1, pattern) != len(str1)/len(pattern) ||
		strings.Count(str2, pattern) != len(str2)/len(pattern) {
		return ""
	}

	return pattern
}
