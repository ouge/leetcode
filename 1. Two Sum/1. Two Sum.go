package leetcode

func twoSum(nums []int, target int) []int {
	posMap := map[int]int{}
	for i := range nums {
		pos, ok := posMap[target-nums[i]]
		if !ok {
			posMap[nums[i]] = i
			continue
		}
		return []int{pos, i}
	}
	return nil
}
