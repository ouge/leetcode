package leetcode

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m <= 1 || n <= 1 {
		return 1
	}
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}
	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[m-1]
}
