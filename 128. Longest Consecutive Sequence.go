package leetcode

func longestConsecutive(nums []int) (ret int) {
	begMap := make(map[int]int)
	endMap := make(map[int]int)
	exitMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := exitMap[num]; ok {
			continue
		}
		exitMap[num] = 0
		end, ok1 := begMap[num+1]
		beg, ok2 := endMap[num-1]

		if ok1 && ok2 {
			// num将两个序列连起来
			begMap[beg] = end
			endMap[end] = beg
			delete(begMap, num+1)
			delete(endMap, num-1)
		} else if ok1 {
			// num与 num+1~... 连接
			begMap[num] = end
			endMap[end] = num
			delete(begMap, num+1)
		} else if ok2 {
			// num与 ...~num-1 连接
			begMap[beg] = num
			endMap[num] = beg
			delete(endMap, num-1)
		} else {
			begMap[num] = num
			endMap[num] = num
		}
	}
	for v1, v2 := range begMap {
		if ret < v2-v1+1 {
			ret = v2 - v1 + 1
		}
	}
	return

}
