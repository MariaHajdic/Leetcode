package main

func gameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			neighbours := 0

			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}
					ni, nj := i+di, j+dj
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
						if board[ni][nj] == 1 || board[ni][nj] == -1 {
							neighbours++
						}
					}
				}
			}

			if board[i][j] == 1 && (neighbours < 2 || neighbours > 3) {
				board[i][j] = -1
			}
			if board[i][j] == 0 && neighbours == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}
