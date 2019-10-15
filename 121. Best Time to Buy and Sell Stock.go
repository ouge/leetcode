package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var minPrice = prices[0]
	var ret int
	for i := 1; i < len(prices); i++ {
		if ret < prices[i]-minPrice {
			ret = prices[i] - minPrice
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return ret
}
