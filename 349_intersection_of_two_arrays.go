package main

func intersection(nums1 []int, nums2 []int) []int {
	unique := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		if _, ok := unique[nums1[i]]; !ok {
			unique[nums1[i]] = false
		}
	}

	result := []int{}
	for i := 0; i < len(nums2); i++ {
		val, ok := unique[nums2[i]]
		if ok && !val { // unique value found
			result = append(result, nums2[i])
			unique[nums2[i]] = true
		}
	}
	return result
}
