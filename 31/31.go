package leetcode

import "sort"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[len(nums)-1] {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
			sort.Sort(sort.IntSlice(nums[i+1:]))
			return
		}
	}
	sort.Reverse(sort.IntSlice(nums))
}
