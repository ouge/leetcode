package leetcode

// 复杂度 mn
func trap(height []int) (ret int) {
	posMap := make(map[int]int) // 前一个相同位置下标
	for i := 0; i < len(height); i++ {
		for j := 1; j <= height[i]; j++ {
			if last, ok := posMap[j]; ok {
				ret += i - last - 1
			}
			posMap[j] = i
		}
	}
	return
}
