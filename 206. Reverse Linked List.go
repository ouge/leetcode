package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var i, j, k *ListNode
	i = head
	j = head.Next
	i.Next = nil

	for j != nil {
		k = j.Next
		j.Next = i
		i = j
		j = k
	}
	return i
}
