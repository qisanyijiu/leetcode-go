package medium

import "testing"

func TestSortList(t *testing.T) {
	l := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	if !sortList(l).equal(&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}){
		t.Fail()
	}
	l = &ListNode{
		Val:  -1,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}
	if !sortList(l).equal(&ListNode{
		Val:  -1,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}){
		t.Fail()
	}
}