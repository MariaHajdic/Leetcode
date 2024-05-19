package main

import "sort"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	alice_sum, bob_sum := 0, 0
	for _, candy := range aliceSizes {
		alice_sum += candy
	}
	for _, candy := range bobSizes {
		bob_sum += candy
	}

	// total_candies /= 2 // A = 10, needs 13 -> goal = 3
	delta := (alice_sum - bob_sum) / 2
	sort.Ints(bobSizes)

	// Sharing candies via binary search
	for _, x := range aliceSizes {
		y := x - delta
		if binarySearch(bobSizes, y) {
			return []int{x, y}
		}
	}

	return nil
}

func binarySearch(arr []int, goal int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == goal {
			return true
		}
		if arr[mid] < goal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

/* Without binary search */ /*
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    aliceSum, bobSum := 0, 0

    for _, candy := range aliceSizes {
        aliceSum += candy
    }
    bobSet := make(map[int]bool)
    for _, candy := range bobSizes {
        bobSum += candy
        bobSet[candy] = true
    }

    delta := (aliceSum - bobSum) / 2

    for _, x := range aliceSizes {
        if bobSet[x-delta] {
            return []int{x, x-delta}
        }
    }

    return nil
}*/
