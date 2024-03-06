package main

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
			continue
		} else if nums[i] == candidate {
			count += 1
		} else {
			count -= 1
		}
		if count >= len(nums)/2+1 {
			return candidate
		}
	}

	return candidate

	// frequencyCounter := make(map[int]int)
	// max := 0
	// for _, num := range nums {
	//     frequencyCounter[num] += 1
	//     if frequencyCounter[num] >= len(nums) / 2  + 1 {
	//         return num
	//     }
	// }
	// return max
	// max := math.MinInt32
	// var res int
	// for key, value := range frequencyCounter {
	//     if value > max {
	//         max = value
	//         res = key
	//     }
	// }
	// return res
}
