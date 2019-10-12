package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nums1, nums2 []int
	head := &ListNode{}
	p := head

	for i := l1; i != nil; i = i.Next {
		nums1 = append(nums1, i.Val)
	}
	for i := l2; i != nil; i = i.Next {
		nums2 = append(nums2, i.Val)
	}
	var carry int
	var i, j int
	for i, j = len(nums1)-1, len(nums2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		v := nums1[i] + nums2[j] + carry
		if v >= 10 {
			carry = 1
			v -= 10
		} else {
			carry = 0
		}
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	if j >= 0 {
		i = j
		nums1 = nums2
	}
	for ; i >= 0; i-- {
		v := nums1[i] + carry
		if v >= 10 {
			carry = 1
			v -= 10
		} else {
			carry = 0
		}
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	if carry == 1 {
		p.Next = &ListNode{Val: 1}
		p = p.Next
	}
	p = head.Next
	if p == nil {
		return nil
	}
	q := p.Next
	p.Next = nil
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}

	return p
}
