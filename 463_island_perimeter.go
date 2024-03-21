package main

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				perimeter += 4

				if i > 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}

	return perimeter
}
