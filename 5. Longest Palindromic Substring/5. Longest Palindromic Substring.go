package leetcode

// TODO: 优化算法
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var ret string
	res := make([][]int, len(s))
	ret = s[:1]
	for i := range res {
		res[i] = make([]int, len(s))
		res[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}
			if j == i+1 {
				res[i][j] = 2
				if res[i][j] > len(ret) {
					ret = s[i : j+1]
				}
				continue
			}
			if res[i+1][j-1] == 0 {
				continue
			}
			res[i][j] = res[i+1][j-1] + 2
			if res[i][j] > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}
