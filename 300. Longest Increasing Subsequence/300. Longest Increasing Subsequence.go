package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minTail := make([]int, len(nums)+1) // 长度为k的LIS最小尾部为minTail[k]
	minTail[1] = nums[0]
	n := 1
	var beg, end int

	for i := 1; i < len(nums); i++ {
		// 在minTail[1...n]中找到最小的>=nums[i]的位置
		beg, end = 1, n+1
		for beg < end {
			mid := (beg + end) / 2
			if minTail[mid] < nums[i] {
				beg = mid + 1
			} else {
				end = mid
			}
		}
		minTail[beg] = nums[i]
		if beg == n+1 {
			n++
		}
	}
	return n
}
