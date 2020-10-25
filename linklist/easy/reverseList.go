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