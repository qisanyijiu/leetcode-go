package easy

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T){
	l1 := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(&l1, &l2)
	arr := []int{}
	for result != nil{
		arr = append(arr, result.Val)
		result = result.Next
	}
	expect := []int{1, 1, 2, 3, 4, 4}
	for i:=0; i< len(arr); i++ {
		if arr[i] != expect[i]{
			t.Fail()
		}
	}
}