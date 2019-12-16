package leetcode

func grayCode(n int) []int {
	ret := make([]int, 1<<n)
	for i := 0; i < n; i++ {
		for j, k := 0, (1<<(i+1))-1; j < k; j, k = j+1, k-1 {
			ret[k] = ret[j] + (1 << i)
		}
	}
	return ret
}
