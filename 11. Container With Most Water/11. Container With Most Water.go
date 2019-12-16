package leetcode

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	var res int
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			res = maxInt(res, (j-i)*height[i])
			for {
				i++
				if i == j || (i < j && height[i] > height[i-1]) {
					break
				}
			}
		} else {
			res = maxInt(res, (j-i)*height[j])
			for {
				j--
				if i == j || (i < j && height[j] > height[j+1]) {
					break
				}
			}
		}
	}
	return res
}
