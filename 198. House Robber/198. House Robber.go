package leetcode

func rob(nums []int) int {
	var a, b int
	for i := range nums {
		if b < a+nums[i] {
			a, b = b, a+nums[i]
		} else {
			a = b
		}
	}
	return b
}
