package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	result := [][]int{}
	curr := []int{}

	for i := 0; i < len(intervals); i++ {
		curr = intervals[i]
		if curr[1] < newInterval[0] {
			result = append(result, curr)
			continue
		}
		if curr[0] > newInterval[1] {
			return append(append(result, newInterval), intervals[i:]...)
		}
		if curr[0] < newInterval[0] {
			newInterval[0] = curr[0]
		}
		if curr[1] <= newInterval[1] {
			continue
		}
		newInterval[1] = curr[1]
	}
	return append(result, newInterval)
}
