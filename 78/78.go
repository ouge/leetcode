package leetcode

func subsets(nums []int) (ret [][]int) {
	if len(nums) == 1 {
		return [][]int{{nums[0]}, {}}
	}
	ret = subsets(nums[:len(nums)-1])
	n := len(ret)
	for i := 0; i < n; i++ {
		newSub := append([]int{}, ret[i]...)
		ret = append(ret, append(newSub, nums[len(nums)-1]))
	}
	return ret
}
