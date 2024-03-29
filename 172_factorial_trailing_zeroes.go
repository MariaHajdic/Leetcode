package main

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	res := 0
	for maxPower := 5; maxPower <= n; maxPower *= 5 {
		for i := maxPower; i <= n; i += maxPower {
			res++
		}
	}
	return res
}
