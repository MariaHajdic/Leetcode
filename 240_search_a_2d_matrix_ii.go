package main

func searchMatrix_ii(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for row, col := 0, n-1; row < m && col >= 0; {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col-- // the entire last col is greater than target
		} else {
			row++ // the entire row is smaller than target
		}
	}

	return false
}
