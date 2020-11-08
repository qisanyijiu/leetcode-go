package easy

import "testing"

func TestKthToLast(t *testing.T) {
	list := &ListNode{
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
	if res := KthToLast(list, 2); res != 4 {
		t.Errorf("except: %d, got: %d", 4, res)
	}

}
