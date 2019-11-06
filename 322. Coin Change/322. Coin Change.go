package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for i := range coins {
		for j := 1; j <= amount; j++ {
			if j >= coins[i] && dp[j-coins[i]] != -1 {
				if dp[j] == -1 || dp[j] > dp[j-coins[i]]+1 {
					dp[j] = dp[j-coins[i]] + 1
				}
			}
		}
	}
	return dp[amount]
}
