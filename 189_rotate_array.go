package main

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/* Additional O(n) space.

func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n // Might be greater than n.

    rotatedArray := make([]int, n)
    for i := range n {
        rotatedArray[(i+k)%n] = nums[i]
    }
    copy(nums, rotatedArray)
}
*/
