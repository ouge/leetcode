package leetcode

func search(nums []int, target int) int {

	for beg, end := 0, len(nums)-1; beg <= end; {
		mid := beg + (end-beg)/2
		if nums[mid] == target {
			return mid
		}

		if nums[beg] <= nums[mid] {
			// mid在左侧递增序列
			if nums[beg] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				beg = mid + 1
			}
		} else {
			// mid在右侧递增序列
			if nums[mid] < target && target <= nums[end] {
				beg = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1

}
