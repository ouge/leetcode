package leetcode

func permute(nums []int) [][]int {
	res := doPermute(nums, 0, nil)
	return res
}

func doPermute(nums []int, k int, res [][]int) [][]int {
	if k == len(nums)-1 {
		newOne := make([]int, len(nums))
		copy(newOne, nums)
		res = append(res, newOne)
		return res
	}

	for i := k; i < len(nums); i++ {
		nums[i], nums[k] = nums[k], nums[i]
		res = doPermute(nums, k+1, res)
		nums[i], nums[k] = nums[k], nums[i]
	}
	return res
}
