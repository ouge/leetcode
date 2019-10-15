package leetcode

func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return nil
	}
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for i := 0; len(nodes) > 0; i++ {
		var newNodes []*TreeNode
		var nums []int
		for i := range nodes {
			nums = append(nums, nodes[i].Val)
			if nodes[i].Left != nil {
				newNodes = append(newNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				newNodes = append(newNodes, nodes[i].Right)
			}
		}
		if i&1 == 1 {
			// i为偶数时翻转
			for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nodes = newNodes
		ret = append(ret, nums)
	}
	return ret
}
