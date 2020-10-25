package easy

import "testing"

func TestRemoveElements(t *testing.T){
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  6,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	if !removeElements(l, 6).equal(&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}){
		t.Fail()
	}
}
