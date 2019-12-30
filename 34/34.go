package leetcode

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	beg, end := 0, len(nums)-1
	for beg < end {
		mid := beg + (end-beg)/2
		if nums[mid] < target {
			beg = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			end = mid
		}
	}
	if nums[beg] != target {
		return ret
	}
	ret[0] = beg
	beg, end = 0, len(nums)-1
	for beg < end {
		mid := beg + (end-beg+1)/2
		if nums[mid] < target {
			beg = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			beg = mid
		}
	}
	ret[1] = beg
	return ret
}
