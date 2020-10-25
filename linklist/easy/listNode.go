package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode)length()int{
	cnt := 0
	for l != nil{
		cnt += 1
		l = l.Next
	}
	return cnt
}

func (n *ListNode)equal(l *ListNode)bool{
	for n != nil && l != nil{
		if n.Val == l.Val{
			n=n.Next
			l=l.Next
		}else{
			return false
		}
	}
	if n == nil && l != nil{
		return false
	}else if n != nil && l == nil{
		return false
	}else{
		return true
	}
}
