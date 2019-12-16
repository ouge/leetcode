package leetcode

func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}
	ret := make([]int, len(nums))

	ret[0] = 1
	for i := 1; i < len(ret); i++ {
		ret[i] = nums[i-1] * ret[i-1]
	}
	t := 1
	for i := len(ret) - 1; i >= 0; i-- {
		ret[i] = ret[i] * t
		t *= nums[i]
	}
	return ret
}
