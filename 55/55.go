package leetcode

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func canJump(nums []int) bool {
	ret := make([]bool, len(nums))
	ret[0] = true
	for i := 0; i < len(nums) && ret[i]; i++ {
		for j := min(i+nums[i], len(nums)-1); j >= i && !ret[j]; j-- {
			ret[j] = true
		}
	}
	return ret[len(ret)-1]
}
