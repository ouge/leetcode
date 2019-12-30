package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hhead := &ListNode{Next: head}
	p := hhead
	for n >= 0 {
		p = p.Next
		n--
	}
	q := hhead
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return hhead.Next
}
