package leetcode

import (
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func threeSumClosest(nums []int, target int) (ret int) {
	ret = nums[0] + nums[1] + nums[2]
	dif := abs(target - ret)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		t := target - nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			r := nums[j] + nums[k]
			if abs(t-r) < dif {
				ret = r + nums[i]
				dif = abs(t - r)
			}
			if r < t {
				j++
			} else if r > t {
				k--
			} else {
				return target
			}
		}
	}
	return
}
