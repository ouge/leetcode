package leetcode

func trap(height []int) (ret int) {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	maxHeight := 0
	for i := 0; i < len(height); i++ {
		leftMax[i] = maxHeight
		if height[i] >= maxHeight {
			maxHeight = height[i]
		}
	}
	maxHeight = 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMax[i] = maxHeight
		if height[i] >= maxHeight {
			maxHeight = height[i]
		}
	}

	for i := range height {
		minHeight := leftMax[i]
		if minHeight > rightMax[i] {
			minHeight = rightMax[i]
		}
		if minHeight > height[i] {
			ret += minHeight - height[i]
		}
	}

	return
}
