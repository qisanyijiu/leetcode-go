package easy

func isPalindrome(head *ListNode) bool {
	if reverseList(head).equal(head){
		return true
	}
	return false
}