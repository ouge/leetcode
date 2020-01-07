package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	beg, end := 0, len(lists)
	mid := beg + (end-beg)/2

	left, right := mergeKLists(lists[beg:mid]), mergeKLists(lists[mid:])

	res := &ListNode{}
	p := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	for left != nil {
		p.Next = left
		left = left.Next
		p = p.Next
	}
	for right != nil {
		p.Next = right
		right = right.Next
		p = p.Next
	}

	return res.Next
}
