package leetcode

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}
