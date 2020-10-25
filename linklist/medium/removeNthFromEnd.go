package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := head.length()
	if length == 1 {
		return nil
	}
	if length == n {
		return head.Next
	}
	p := 0
	out := head
	for {
		p ++
		if p == length - n{
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return out
}