package easy

import "testing"

func TestMiddleNode(t *testing.T) {
	l := &ListNode{
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
	}
	r := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	if !middleNode(l).equal(r){
		t.Fail()
	}
	l = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
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
	}
	r = &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	if !middleNode(l).equal(r){
		t.Fail()
	}
}
