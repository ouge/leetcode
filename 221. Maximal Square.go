package leetcode

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	maxSquare := make([][]int, n)
	maxX := make([][]int, n)
	maxY := make([][]int, n)
	maxZ := make([][]int, n)
	for i := range maxSquare {
		maxSquare[i] = make([]int, m)
		maxX[i] = make([]int, m)
		maxY[i] = make([]int, m)
		maxZ[i] = make([]int, m)
	}

	if matrix[0][0] == 1 {
		maxSquare[0][0] = 1
		maxX[0][0] = 1
		maxY[0][0] = 1
		maxZ[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if matrix[i][0] == 1 {
			maxSquare[i][0] = 1
			maxX[i][0] = 1
			maxY[i][0] = 1
			maxZ[i][0] = 1
		} else {
			maxSquare[i][0] = matrix[i-1][0]
			maxY[i][0] = 0
			maxX[i][0] =
		}
	}

	for i := 1; i < n; i++ {
		if matrix[i][0] == 1 {

		}

		for j := 1; j < m; j++ {
			var v int
			if i > 0 && v < maxSquare[i-1][j] {
				v = maxSquare[i-1][j]
			}
			if j > 0 && v < maxSquare[i][j-1] {
				v = maxSquare[i][j-1]
			}
			if matrix[i][j] == 1 && {

			}

		}
	}

}
