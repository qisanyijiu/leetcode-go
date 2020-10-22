package easy

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var out = &ListNode{}
	head := out
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			out.Next = l1
			l1 = l1.Next
			out = out.Next
		} else {
			out.Next = l2
			l2 = l2.Next
			out = out.Next
		}
	}
	if l1 != nil{
		out.Next = l1
	}else if l2 != nil{
		out.Next = l2
	}
	return head.Next
}
