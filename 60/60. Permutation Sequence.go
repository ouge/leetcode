package leetcode

// n2 复杂度
func getPermutation(n int, k int) string {
	use := make([]int, n+1)       // 数字是否使用，1为已用
	factorial := make([]int, n+1) // 阶乘结果, factorial[x] = x!
	factorial[0] = 1
	var res []byte

	// 计算阶乘
	for i := 1; i <= n; i++ {
		factorial[i] = i * factorial[i-1]
	}

	var pos int

	k--
	for i := 0; i < n; i++ {
		pos = k / factorial[n-1-i] // 剩余元素的相对位置
		k = k % factorial[n-1-i]

		cnt := 0
		for j := 1; j <= n; j++ {
			if use[j] == 0 { // 剩余
				if cnt >= pos {
					res = append(res, '0'+byte(j))
					use[j] = 1
					break
				}
				cnt++
			}
		}
	}

	return string(res)
}
