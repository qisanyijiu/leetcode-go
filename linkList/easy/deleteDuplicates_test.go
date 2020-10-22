package easy

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	arr := []int{}
	data := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	result := deleteDuplicates(data)
	arr = []int{}
	for result != nil{
		arr = append(arr, result.Val)
		result = result.Next
	}
	expect := []int{1,2}
	for i, val := range arr{
		if val != expect[i]{
			t.Fail()
		}
	}
	data = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	result = deleteDuplicates(data)
	arr = []int{}
	for result != nil{
		arr = append(arr, result.Val)
		result = result.Next
	}
	expect = []int{1,2,3}
	for i, val := range arr{
		if val != expect[i]{
			t.Fail()
		}
	}
}