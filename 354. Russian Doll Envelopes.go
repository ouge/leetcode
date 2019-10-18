package leetcode

import "sort"

// 用于对输入排序
>>>>>>> origin/master
type sortEnvelopes [][]int

func (s sortEnvelopes) Len() int {
	return len(s)
}
func (s sortEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] > s[j][1]
	}
	return s[i][0] < s[j][0]
}
func (s sortEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(sortEnvelopes(envelopes))

	minTail := make([]int, len(envelopes)+1) // min[i]: 递增子序列长度为i的最小尾部
	n := 0
	var beg, end int
	for i := range envelopes {
		// 二分查找第一个比 envelopes[i][1] 大的位置
		for beg, end = 1, n+1; beg < end; {
			mid := beg + (end-beg)/2
			if minTail[mid] < envelopes[i][1] {
				beg = mid + 1
			} else {
				end = mid
			}
		}
		minTail[beg] = envelopes[i][1]
		if beg > n {
			n = beg
		}
	}
	return n
}
