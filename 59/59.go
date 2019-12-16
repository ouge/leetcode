package leetcode

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	min, max := 0, n-1
	val := 1

	for min < max {
		for j := min; j <= max; j++ {
			ret[min][j] = val
			val++
		}
		for i := min + 1; i <= max; i++ {
			ret[i][max] = val
			val++
		}
		for j := max - 1; j >= min; j-- {
			ret[max][j] = val
			val++
		}
		for i := max - 1; i >= min+1; i-- {
			ret[i][min] = val
			val++
		}
		min++
		max--
	}
	if min == max {
		ret[min][min] = val
	}
	return ret
}
