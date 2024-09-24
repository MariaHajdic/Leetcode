package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	idx_stack := []int{} // for indexes of not updated Ts

	for i := 0; i < len(temperatures); i++ {
		// Updating previous Ts while the current one is greater
		for len(idx_stack) > 0 && temperatures[i] > temperatures[idx_stack[len(idx_stack)-1]] {
			idx := idx_stack[len(idx_stack)-1]
			res[idx] = i - idx
			idx_stack = idx_stack[:len(idx_stack)-1]
		}
		idx_stack = append(idx_stack, i)
	}

	return res
}

// [73,74,75,71,69,72,76,73]
// [1,0,0,0,0,0,0,0] res
// [ ] stack
