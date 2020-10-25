package easy


func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	newList := &ListNode{}
	out := newList
	for head != nil{
		if head.Val != val{
			newList.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			newList = newList.Next
		}
		head = head.Next
	}
	return out.Next
}