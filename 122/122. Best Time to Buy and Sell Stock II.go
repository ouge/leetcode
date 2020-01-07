package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profits := make([]int, len(prices))
	maxXXX := - prices[0] // maxXXX[i] = max(profits[k-1] - price[k])
	profits[0] = 0

	for i := 1; i < len(prices); i++ {
		if profits[i-1] > maxXXX+prices[i] {
			profits[i] = profits[i-1]
		} else {
			profits[i] = maxXXX + prices[i]
		}

		if maxXXX < profits[i-1]-prices[i] {
			maxXXX = profits[i-1] - prices[i]
		}
	}
	return profits[len(prices)-1]
}

func maxProfit(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}
