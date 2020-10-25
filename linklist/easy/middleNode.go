package easy

func middleNode(head *ListNode) *ListNode {
	length := head.length()
	half := length / 2
	for i := half; i > 0; i -- {
		head = head.Next
	}
	return head
}