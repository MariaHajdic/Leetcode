package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	freq_map := map[int]int{}
	for _, num := range nums {
		freq_map[num]++
	}

	freq_arr := make([][]int, len(freq_map)) // for each unique number
	i := 0
	for num := range freq_map {
		freq_arr[i] = make([]int, 2)
		freq_arr[i][0] = freq_map[num]
		freq_arr[i][1] = num
		i++
	}

	// the map values are sorted by value now, i.e. desc by frequency
	sort_arr(freq_arr)

	res := make([]int, k)
	for i = 0; i < k; i++ {
		res[i] = freq_arr[i][1]
	}

	return res
}

func sort_arr(arr [][]int) {
	sort.Slice(arr[:], func(i, j int) bool {
		for x := range arr[i] {
			return arr[i][x] > arr[j][x] // descending order
		}
		return false
	})
}
