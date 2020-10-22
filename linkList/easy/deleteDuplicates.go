package easy

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	newList := &ListNode{
		Val:head.Val,
	}
	out := newList
	head = head.Next
	for head != nil{
		if head.Val != newList.Val {
			newList.Next = head
			newList = newList.Next
		}
		head = head.Next
	}
	newList.Next = nil
	return out
}