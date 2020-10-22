package easy

import "testing"

func TestHasCycle(t *testing.T){
	var data *ListNode
	data = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	item := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  -4,
			Next: data,
		},
	}
	data.Next.Next = item
	if !hasCycle(data) {
		t.Fail()
	}
	data = &ListNode{
		Val:  1,
		Next: nil,
	}
	item = &ListNode{
		Val:  2,
		Next: data,
	}
	data.Next = item
	if !hasCycle(data) {
		t.Fail()
	}

	data = &ListNode{Val:1}
	if hasCycle(data) {
		t.Fail()
	}
}