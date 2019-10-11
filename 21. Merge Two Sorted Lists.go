package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	retHead := &ListNode{}
	ptr := retHead

	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		ptr.Next = node
		ptr = ptr.Next
	}
	if l1 != nil {
		ptr.Next = l1
	} else if l2 != nil {
		ptr.Next = l2
	}

	return retHead.Next
}
