package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	ranks := make([][2]int, len(score))

	for i, s := range score {
		ranks[i][0] = s // score
		ranks[i][1] = i // original idx
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i][0] > ranks[j][0]
	})

	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	res := make([]string, len(score))

	for i := 0; i < len(ranks); i++ {
		idx := ranks[i][1] // get original index
		if i < 3 {
			res[idx] = medals[i]
		} else {
			res[idx] = strconv.Itoa(i + 1)
		}
	}

	return res
}

// func findRelativeRanks(score []int) []string {
//     h := &MaxHeap{}
//     heap.Init(h)
//     n := len(score)

//     for i, s := range score {
//         heap.Push(h, Item{score: s, index: i})
//     }

//     result := make([]string, n)
//     medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

//     for rank := 1; h.Len() > 0; rank++ {
//         item := heap.Pop(h).(Item)
//         if rank <= 3 {
//             result[item.index] = medals[rank-1]
//         } else {
//             result[item.index] = strconv.Itoa(rank)
//         }
//     }

//     return result
// }

// type Item struct {
//     score int
//     index int
// }

// type MaxHeap []int

// func (h MaxHeap) Len() int { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i].score > h[j].score }
// func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
// func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
// func (h *MaxHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[:n-1]
//     return x
// }
