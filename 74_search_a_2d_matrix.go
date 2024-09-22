package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][n-1] < target {
			continue
		}
		return binarySearch_(matrix[i], target)
	}
	return false
}

func binarySearch_(arr []int, target int) bool {
	for left, right := 0, len(arr)-1; left <= right; {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
			continue
		}
		right = mid - 1
	}

	return false
}
