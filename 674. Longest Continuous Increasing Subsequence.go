package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 1
	maxLen := 1 // 以当前元素为序列末尾的最长递增序列长度

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			maxLen++
		} else {
			maxLen = 1
		}

		if ret < maxLen {
			ret = maxLen
		}
	}
	return ret
}
