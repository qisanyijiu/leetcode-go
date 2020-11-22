package easy

import "testing"

func TestReverseList(t *testing.T){
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	if !reverseList(l).equal(&ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}){
		t.Fail()
	}
}


func TestReverseListA(t *testing.T){
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	if !ReverseList(l).equal(&ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}){
		t.Fail()
	}
}
