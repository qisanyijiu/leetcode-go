package easy

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T){
	la := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  8,
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
	lb := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  8,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	if !getIntersectionNode(la, lb).equal(&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  8,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}){
		t.Fail()
	}


	la = &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	lb = &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	if !getIntersectionNode(la, lb).equal(&ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}){
		t.Fail()
	}

	la = &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	lb = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	if getIntersectionNode(la, lb) != nil {
		t.Fail()
	}
}