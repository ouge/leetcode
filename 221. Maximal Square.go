package leetcode

func maximalSquare(matrix [][]byte) (ret int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	maxSquare := make([][]int, n) // maxSquare[i][j]: 以i,j为正方形右下角的最大正方形边长
	for i := 0; i < n; i++ {
		maxSquare[i] = make([]int, m)
		if matrix[i][0] == '1' {
			maxSquare[i][0] = 1
			ret = 1
		}
	}
	for j := 0; j < m; j++ {
		if matrix[0][j] == '1' {
			maxSquare[0][j] = 1
			ret = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				// 只有当该块是1时，才需要计算，否则边长为0
				// maxSquare[i][j] = min(maxSquare[i-1][j], maxSquare[i-1][j-1], maxSquare[i][j-1]) + 1
				maxSquare[i][j] = maxSquare[i-1][j-1]
				if maxSquare[i][j] > maxSquare[i-1][j] {
					maxSquare[i][j] = maxSquare[i-1][j]
				}
				if maxSquare[i][j] > maxSquare[i][j-1] {
					maxSquare[i][j] = maxSquare[i][j-1]
				}
				maxSquare[i][j]++
				if ret < maxSquare[i][j] {
					ret = maxSquare[i][j]
				}
			}
		}
	}
	return ret * ret
}
