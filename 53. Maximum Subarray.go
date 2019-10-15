package leetcode

func maxSubArray(nums []int) (ret int) {
	if len(nums) == 0 {
		return 0
	}
	maxTail := nums[0]
	ret = nums[0]
	for i := 1; i < len(nums); i++ {
		if maxTail <= 0 {
			maxTail = nums[i]
		} else {
			maxTail += nums[i]
		}

		if ret < maxTail {
			ret = maxTail
		}
	}
	return
}
