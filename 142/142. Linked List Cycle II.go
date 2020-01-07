package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	if slow != fast {
		return nil
	}

	third := head
	slow = slow.Next
	for third != slow {
		slow, third = slow.Next, third.Next
	}
	return third
}
