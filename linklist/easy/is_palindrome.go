package easy


/**
234. 回文链表

https://leetcode-cn.com/problems/palindrome-linked-list/

请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true

 */
func IsPalindrome(head *ListNode) bool {
	return isPalindromeA(head)
}

func isPalindrome(head *ListNode) bool {
	if reverseList(head).equal(head){
		return true
	}
	return false
}

func isPalindromeA(head *ListNode) bool {
	var slow *ListNode = head
	var fast *ListNode = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	var left *ListNode = head
	var right *ListNode = ReverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left =left.Next
		right = right.Next
	}
	return true
}