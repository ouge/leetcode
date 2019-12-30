package leetcode

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		beg, end := i+1, len(nums)-1
		for beg < end {
			mid := beg + (end-beg+1)/2
			if nums[mid] <= nums[i] {
				end = mid - 1
			} else {
				beg = mid
			}
		}
		nums[beg], nums[i] = nums[i], nums[beg]
	}
	for beg, end := i+1, len(nums)-1; beg < end; beg, end = beg+1, end-1 {
		nums[beg], nums[end] = nums[end], nums[beg]
	}

}
