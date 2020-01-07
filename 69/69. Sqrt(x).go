package leetcode

func mySqrt(x int) int {
	beg, end := 0, x
	var mid int
	for beg < end {
		mid = beg + (end-beg+1)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid <= x {
			beg = mid
		} else {
			end = mid - 1
		}
	}
	return beg
}
