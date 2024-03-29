package main

func largestAltitude(gain []int) int {
	current, highest := 0, 0
	for _, netGain := range gain {
		current += netGain
		if current > highest {
			highest = current
		}
	}
	return highest
}
