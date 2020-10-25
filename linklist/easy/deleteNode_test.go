package easy

import "testing"

func TestDeleteNode(t *testing.T) {
	l := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	deleteNode(l)
}
