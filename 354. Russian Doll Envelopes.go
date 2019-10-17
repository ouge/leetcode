package leetcode

import "sort"

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



	return 0
}
