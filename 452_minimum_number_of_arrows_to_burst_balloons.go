package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := 1
	prev, curr := points[0], []int{}
	for i := 1; i < len(points); i++ {
		curr = points[i]
		if curr[0] <= prev[1] {
			prev[0] = curr[0]
			if curr[1] < prev[1] {
				prev[1] = curr[1]
			}
			continue
		}
		prev = curr
		result++
	}
	return result
}
