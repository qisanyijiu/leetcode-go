package hard

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var a *ListNode = head
	var b *ListNode = head
	for i := 0 ; i < k; i ++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := Reverse(a, b)
	a.Next = ReverseKGroup(b, k)
	return newHead

}

func Reverse(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = a
	var next *ListNode = a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}