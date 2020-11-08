package easy


func KthToLast(head *ListNode, k int) int {
	right := head
	for i := 0; i < k;i ++ {
		right = right.Next
	}
	for right != nil {
		head = head.Next
		right = right.Next
	}
	return head.Val
}