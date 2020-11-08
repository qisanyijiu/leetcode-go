package easy

import "strings"

/**
459. 重复的子字符串

https://leetcode-cn.com/problems/repeated-substring-pattern/

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"

输出: False
示例 3:

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 */
func repeatedSubstringPattern(s string) bool {
	/**
	假设字符串由n个字串构成，则拼接后的字串为2n个，掐头去尾之后为2n-2个，
	如果此时的字符串至少包含一个原字符串,则说明至少包含n个子串,则2n-2>=n,n>=2.则说明该字符串是周期性结构,最少由两个子串构成.
	如果一个都不包含,即不包含n个子串,则说明2n-2<n,n<2,即n为1,也就是不符合周期性结构
	*/
	if len(s) == 0{
		return false
	}
	size := len(s)

	ss := (s+s)[1: size * 2  - 1]
	return strings.Contains(ss, s)
}

