package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // -4, -1, -1, 0, 1, 2
	results := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // skipping duplicates
			continue
		}

		triplets := twoSum(nums[i+1:], -nums[i]) // array of new solutions
		results = append(results, triplets)
	}

	return results
}

func twoSum_(nums []int, target int) [][]int {
	results := [][]int{}

	for left, right := 0, len(nums)-1; left < right; {
		if nums[left]+nums[right] == target {
			results = append(results, []int{-target, nums[left], nums[right]})
			for left < right && nums[left] == nums[left+1] { // skipping duplicates
				left++
			}
			for left < right && nums[right] == nums[right-1] {
				right--
			}
			left++
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}

	return results
}
