package leetcode

func spiralOrder(matrix [][]int) (ret []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	iMin, iMax := 0, len(matrix)-1
	jMin, jMax := 0, len(matrix[0])-1

	for iMin <= iMax && jMin <= jMax {
		for j := jMin; j <= jMax; j++ {
			ret = append(ret, matrix[iMin][j])
		}
		for i := iMin + 1; i <= iMax; i++ {
			ret = append(ret, matrix[i][jMax])
		}
		if iMin == iMax || jMin == jMax {
			return
		}

		for j := jMax - 1; j >= jMin; j-- {
			ret = append(ret, matrix[iMax][j])
		}
		for i := iMax - 1; i >= iMin+1; i-- {
			ret = append(ret, matrix[i][jMin])
		}
		iMin++
		iMax--
		jMin++
		jMax--
	}
	return ret
}
