package main

func isValidSudoku(board [][]byte) bool {
	boxes := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		boxes[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			boxIdx := (i/3)*3 + j/3

			if boxes[boxIdx][num] || rows[i][num] || cols[j][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIdx][num] = true
		}
	}

	return true
}
