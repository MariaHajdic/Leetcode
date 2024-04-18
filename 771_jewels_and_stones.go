package main

func numJewelsInStones(jewels string, stones string) int {
	types := map[rune]bool{}
	for _, jewel := range jewels {
		types[jewel] = true
	}

	result := 0
	for _, stone := range stones {
		if types[stone] {
			result++
		}
	}

	return result
}
