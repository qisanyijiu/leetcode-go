package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	if isPalindrome(l){
		t.Fail()
	}
	l = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	if !isPalindrome(l){
		t.Fail()
	}
}
