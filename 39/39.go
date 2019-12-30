package leetcode

func combinationSum(candidates []int, target int) [][]int {
	return dfs(candidates, 0, target, make([]int, 0, target), nil)
}

func dfs(candidates []int, idx int, target int, nums []int, res [][]int) [][]int {
	if target == 0 {
		return append(res, append([]int{}, nums...))
	}
	if idx == len(candidates) {
		return res
	}
	res = dfs(candidates, idx+1, target, nums, res)
	for i := candidates[idx]; i <= target; i += candidates[idx] {
		nums = append(nums, candidates[idx])
		res = dfs(candidates, idx+1, target-i, nums, res)
	}
	return res
}
