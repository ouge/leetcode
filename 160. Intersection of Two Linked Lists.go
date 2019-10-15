package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	var n1, n2 int
	for pa != nil {
		n1++
		pa = pa.Next
	}
	for pb != nil {
		n2++
		pb = pb.Next
	}
	if n1 > n2 {
		for i := 0; i < n1-n2; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < n2-n1; i++ {
			headB = headB.Next
		}
	}
	for headA != nil && headA != headB {
		headA, headB = headA.Next, headB.Next
	}
	return headA

}
