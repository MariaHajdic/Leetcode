package main

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	exists := map[int]bool{}
	for _, num := range nums {
		exists[num] = true
	}

	longest_streak := 1

	for num := range exists {
		if !exists[num-1] {
			current_num, current_streak := num, 1

			for exists[current_num+1] {
				current_num++
				current_streak++
			}

			if current_streak > longest_streak {
				longest_streak = current_streak
			}
		}
	}

	return longest_streak
}
