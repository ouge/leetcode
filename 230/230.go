package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	_, res := nodeNumOrResult(root, k)
	return res
}

func nodeNumOrResult(root *TreeNode, k int) (bool, int) {
	if root == nil {
		return false, 0
	}
	ok1, res1 := nodeNumOrResult(root.Left, k)
	if ok1 {
		return true, res1
	}
	if res1 == k-1 {
		return true, root.Val
	}
	ok2, res2 := nodeNumOrResult(root.Right, k-res1-1)
	if ok2 {
		return true, res2
	}
	return false, res1 + res2 + 1
}
