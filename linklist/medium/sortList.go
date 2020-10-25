package medium

/**
通过快慢指针进行归并排序，🐂🍺！！
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(mid))
}

func merge(one, two *ListNode) *ListNode {
	three := new(ListNode)
	tmp := three
	for one != nil && two != nil {
		if one.Val < two.Val {
			tmp.Next = one
			one = one.Next
			tmp = tmp.Next
		} else {
			tmp.Next = two
			two = two.Next
			tmp = tmp.Next
		}
	}
	if one != nil {
		tmp.Next = one
	}
	if two != nil {
		tmp.Next = two
	}
	return three.Next
}