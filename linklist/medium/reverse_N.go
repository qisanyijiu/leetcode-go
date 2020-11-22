package medium


func ReverseN(head *ListNode, n int) *ListNode {
	var out *ListNode
	out = nil
	var tail *ListNode
	tail = nil
	for head != nil && n > 0 {
		out = &ListNode{
			Val: head.Val,
			Next: out,
		}
		if out.Next == nil {
			tail = out
		}
		head = head.Next
		n --
	}
	if tail != nil && head != nil{
		tail.Next = head
	}
	return out
}