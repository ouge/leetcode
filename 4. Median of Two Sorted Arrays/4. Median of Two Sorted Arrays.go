package leetcode

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		nums1, nums2, n1, n2 = nums2, nums1, n2, n1
	}
	if n1 == 0 {
		return 0
	}

	iMin, iMax, halfLen := 0, n1, (n1+n2+1)/2
	var maxOfLeft, minOfRight int

	for iMin <= iMax {
		i := (iMax + iMin) / 2
		j := halfLen - i
		if i < n1 && nums2[j-1] > nums1[i] {
			iMin = i + 1
			continue
		} else if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
			continue
		}

		if i == 0 {
			maxOfLeft = nums2[j-1]
		} else if j == 0 {
			maxOfLeft = nums1[i-1]
		} else {
			maxOfLeft = maxInt(nums1[i-1], nums2[j-1])
		}

		if (n1+n2)%2 == 1 {
			return float64(maxOfLeft)
		}
		if i == n1 {
			minOfRight = nums2[j]
		} else if j == n2 {
			minOfRight = nums1[i]
		} else {
			minOfRight = minInt(nums1[i], nums2[j])
		}
		return float64(maxOfLeft+minOfRight) / 2
	}
	return 0
}
