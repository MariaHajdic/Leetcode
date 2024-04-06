package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 || n == 1 && len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}

	leftClean := flowerbed[0] == 0

	for i, consecutiveZeroes := 0, 0; i < len(flowerbed); i++ {
		if n <= 0 {
			break
		}
		if flowerbed[i] == 1 {
			diff := 1
			if leftClean {
				diff = 0
			}
			n -= (consecutiveZeroes - diff) / 2
			consecutiveZeroes = 0
			leftClean = false
			continue
		} else if i == len(flowerbed)-1 {
			n -= consecutiveZeroes / 2
		}
		consecutiveZeroes++
	}
	if flowerbed[len(flowerbed)-1] == 0 && len(flowerbed) > 1 && flowerbed[len(flowerbed)-2] == 0 {
		n--
	}

	return n <= 0
}
