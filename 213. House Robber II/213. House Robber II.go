package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var a, b int
	var res int
	for i := 1; i < len(nums); i++ {
		if b < a+nums[i] {
			a, b = b, a+nums[i]
		} else {
			a = b
		}
	}
	a, b, res = 0, 0, b
	for i := 0; i < len(nums)-1; i++ {
		if b < a+nums[i] {
			a, b = b, a+nums[i]
		} else {
			a = b
		}
	}
	if res < b {
		return b
	} else {
		return res
	}
}
