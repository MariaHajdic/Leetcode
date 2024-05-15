package main

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

/* Simple */ /*
func singleNumber(nums []int) int {
    slices.Sort(nums)
    prev := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != prev {
            break
        }
        if i < len(nums) - 1 {
            prev = nums[i+1]
            i++
        }
    }
    return prev
}*/
