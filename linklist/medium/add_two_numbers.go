package medium

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{
		Val:  0,
		Next: nil,
	}
	var cur = res
	var carry int = 0
	for l1 != nil || l2 != nil {
		var (
			x, y = 0, 0
		)
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		var sum int = x + y + carry
		carry = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	if carry == 1 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return res.Next
}
