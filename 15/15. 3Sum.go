package leetcode

import "sort"

func threeSum(nums []int) (ret [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i <= n-3 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sumjk := -nums[i]

		for j := i + 1; j <= n-2 && nums[j] <= sumjk; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sumk := sumjk - nums[j]
			for beg, end := j+1, n-1; beg <= end; {
				mid := beg + ((end - beg) >> 2)
				if nums[mid] > sumk {
					end = mid - 1
				} else if nums[mid] < sumk {
					beg = mid + 1
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[mid]})
					break
				}
			}
		}
	}
	return
}
