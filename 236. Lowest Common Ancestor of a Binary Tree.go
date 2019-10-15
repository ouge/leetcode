package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if p == q {
		return p
	}

	_, ret := isFind(root, p, q)
	return ret
}

func isFind(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	if root == p || root == q {
		// 查到值
		num1, _ := isFind(root.Left, p, q)
		num2, _ := isFind(root.Right, p, q)
		if num1 == 1 || num2 == 1 {
			return 2, root
		} else {
			return 1, nil
		}
	}

	num1, node1 := isFind(root.Left, p, q)
	if num1 == 2 {
		return 2, node1
	}
	if num1 == 0 {
		// 左子树不存在值，查右子树
		return isFind(root.Right, p, q)
	}

	// 左子树有一个值

	num2, _ := isFind(root.Right, p, q)
	if num2 == 1 {
		return 2, root
	}
	return 1, nil
}
