package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 int
	for i := l1; i != nil; i = i.Next {
		num1++
	}
	for i := l2; i != nil; i = i.Next {
		num2++
	}
	if num1 < num2 {
		num1, num2 = num2, num1
		l1, l2 = l2, l1
	}
	cnt := num1 - num2

	var l3, t *ListNode
	for cnt > 0 {
		t = &ListNode{Val: 1}

		l1 = l1.Next

	}

	return nil
}
