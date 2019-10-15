package leetcode

func maximalSquare(matrix [][]byte) int {
	var maxSquare [][]int

	for i := range matrix {
		if len(maxSquare) <= i {
			maxSquare = append(maxSquare, make([]int, len(matrix[i])))
		}
		for j := range matrix[i] {
			var v int
			if i > 0 && v < maxSquare[i-1][j] {
				v = maxSquare[i-1][j]
			}
			if j > 0 && v < maxSquare[i][j-1] {
				v = maxSquare[i][j-1]
			}
			if matrix[i][j] == 1 {

			}

		}
	}

}
