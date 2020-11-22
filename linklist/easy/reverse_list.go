package easy

func reverseList(head *ListNode) *ListNode {
	var out *ListNode
	out = nil
	for head != nil{
		out = &ListNode{
			Val:  head.Val,
			Next: out,
		}
		head = head.Next
	}
	return out
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last *ListNode = ReverseList(head.Next)
	head.Next.Next = head;
	head.Next = nil
	return last
}