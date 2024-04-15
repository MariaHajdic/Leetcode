package main

func findMaxAverage(nums []int, k int) float64 {
	var currentSum float64
	for i := 0; i < k; i++ {
		currentSum += float64(nums[i])
	}
	maxAverage := currentSum / float64(k)

	for i := 0; i < len(nums)-k; i++ {
		currentSum -= float64(nums[i])
		currentSum += float64(nums[i+k])
		if currentSum/float64(k) > maxAverage {
			maxAverage = currentSum / float64(k)
		}
	}
	return maxAverage
}
