package leetcode

import "sort"

func findKthLargest_sb(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest(nums []int, k int) (ret int) {
	// 建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums)-1)
	}
	// 堆排序
	for i := 0; i < k && i < len(nums); i++ {
		ret = nums[0]
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		adjust(nums, 0, len(nums)-i-2)
	}

	return ret
}

func adjust(nums []int, pos int, last int) {
	lc := pos*2 + 1 // 左孩子
	rc := lc + 1    // 右孩子
	if lc > last {
		// 无孩子
		return
	} else if rc > last {
		// 只有左孩子
		if nums[lc] > nums[pos] {
			// 左孩子大
			nums[lc], nums[pos] = nums[pos], nums[lc]
			adjust(nums, lc, last)
		}
	} else {
		// 有左右孩子
		if nums[rc] > nums[pos] && nums[rc] >= nums[lc] {
			// 右孩子最大
			nums[rc], nums[pos] = nums[pos], nums[rc]
			adjust(nums, rc, last)
		} else if nums[lc] > nums[pos] && nums[lc] >= nums[rc] {
			// 左孩子最大
			nums[lc], nums[pos] = nums[pos], nums[lc]
			adjust(nums, lc, last)
		}
	}
}
