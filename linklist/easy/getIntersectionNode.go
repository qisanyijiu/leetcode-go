package easy

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	abs := func(x int) int {
		if x < 0{
			return -x
		}
		return x
	}
	lenA, lenB := headA.length(), headB.length()
	diff := lenA - lenB
	var longer *ListNode
	var shorter *ListNode
	if diff > 0 {
		longer = headA
		shorter = headB
	}else{
		longer = headB
		shorter = headA
	}
	for i := 0; i< abs(diff); i ++ {
		longer = longer.Next
	}
	for longer != nil && shorter != nil{
		if longer.equal(shorter){
			return shorter
		}
		longer = longer.Next
		shorter = shorter.Next
	}
	return nil
}