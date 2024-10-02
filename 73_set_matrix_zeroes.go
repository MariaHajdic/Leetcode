package main

func setZeroes(matrix [][]int) {
	m := len(matrix) // высота
	n := len(matrix[0])

	zeroesRows, zeroesCols := map[int]bool{}, map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroesRows[i] = true
				zeroesCols[j] = true
			}
		}
	}

	for i := 0; i < m; i++ { // rows
		for j := 0; j < n; j++ { // columns
			if zeroesRows[i] == true {
				matrix[i][j] = 0
			} else if zeroesCols[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}
