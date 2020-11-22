package hard

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

func (l *ListNode)ToSlice() []int {
	var arr = []int{}
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

func (n *ListNode)equal(l *ListNode)bool{
	arr1 := n.ToSlice()
	arr2 := l.ToSlice()
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v1 := range arr1 {
		v2 := arr2[i]
		if v1 != v2 {
			return false
		}
	}

	return true
}

