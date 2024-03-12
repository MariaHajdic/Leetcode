package main

import "sort"

func merge_intervals(intervals [][]int) [][]int {
	result := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	prev, curr := intervals[0], intervals[0]

	for i := 1; i < len(intervals); i++ {
		curr = intervals[i]
		if curr[0] <= prev[1] {
			if curr[1] > prev[1] {
				prev[1] = curr[1]
			}
			continue
		}
		result = append(result, prev)
		prev = curr
	}
	result = append(result, prev)
	return result
}
