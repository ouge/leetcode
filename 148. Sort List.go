package leetcode

func sortList(head *ListNode) *ListNode {
	// 插入排序
	retPre := &ListNode{}

	for head != nil {
		q := retPre
		for q.Next != nil && q.Next.Val < head.Val {
			q = q.Next
		}
		if q.Next == nil {
			q.Next = head
			head = head.Next
			q.Next.Next = nil
		} else {
			t := q.Next
			q.Next = head
			head = head.Next
			q.Next.Next = t
		}
	}
	return retPre.Next

}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left, right := sortList(head), sortList(mid)

	ret := &ListNode{}
	h := ret
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	return ret.Next
}
