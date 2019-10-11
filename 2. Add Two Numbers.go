package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := l1
	pre := l1
	for l1 != nil && l2 != nil {
		l1.Val += l2.Val + carry
		if l1.Val >= 10 {
			carry = 1
			l1.Val -= 10
		} else {
			carry = 0
		}
		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		pre.Next = l2
		l1 = l2
	}
	for l1 != nil {
		l1.Val += carry
		if l1.Val >= 10 {
			carry = 1
			l1.Val -= 10
		} else {
			carry = 0
		}
		pre = l1
		l1 = l1.Next
	}
	if carry == 1 {
		pre.Next = &ListNode{Val: 1}
	}

	return head
}
