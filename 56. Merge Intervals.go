package leetcode

import (
	"sort"
)

type sortIntervals [][]int

func (s sortIntervals) Len() int {
	return len(s)
}
func (s sortIntervals) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func (s sortIntervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) (ret [][]int) {
	if len(intervals) == 0 {
		return
	}
	sort.Sort(sortIntervals(intervals))
	ret = append(ret, []int{intervals[0][0], intervals[0][1]})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			// 有重叠
			if intervals[i][1] <= intervals[i-1][1] {
				// 是前一个区间子区间
				intervals[i][0], intervals[i][1] = intervals[i-1][0], intervals[i-1][1]
			} else {
				// 合并两个小区间
				intervals[i][0] = intervals[i-1][0]
				ret[len(ret)-1][0], ret[len(ret)-1][1] = intervals[i][0], intervals[i][1]
			}

		} else {
			ret = append(ret, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return
}
