package medium

func detectCycle(head *ListNode) *ListNode {
	addressArr := []*ListNode{}
	in := func(n *ListNode) bool {
		for _,item := range addressArr{
			if item == n{
				return true
			}
		}
		return false
	}
	for head != nil{
		if !in(head){
			addressArr = append(addressArr, head)
		}else{
			return head
		}
		head = head.Next
	}
	return nil
}