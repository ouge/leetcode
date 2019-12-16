package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var nodeNum int
	for p := head; p != nil; p = p.Next {
		nodeNum++
	}
	k %= nodeNum
	if k == 0 {
		return head
	}
	split := nodeNum - k
	splitNode := head
	for i := 1; i < split; i++ {
		splitNode = splitNode.Next
	}
	last := splitNode.Next
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head
	head = splitNode.Next
	splitNode.Next = nil
	return head

}
