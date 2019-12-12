package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return aaa(root, new(*int))
}

func aaa(root *TreeNode, val **int) bool {
	if root == nil {
		return true
	}
	ok := aaa(root.Left, val)
	if !ok {
		return false
	}
	if *val == nil {
		*val = new(int)
		**val = root.Val
		return aaa(root.Right, val)
	}
	if root.Val <= **val {
		return false
	}
	**val = root.Val
	return aaa(root.Right, val)
}
