package easy

import "strings"

/**
 假设字符串由n个字串构成，则拼接后的字串为2n个，掐头去尾之后为2n-2个，
 如果此时的字符串至少包含一个原字符串,则说明至少包含n个子串,则2n-2>=n,n>=2.则说明该字符串是周期性结构,最少由两个子串构成.
 如果一个都不包含,即不包含n个子串,则说明2n-2<n,n<2,即n为1,也就是不符合周期性结构
 */
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0{
		return false
	}
	size := len(s)

	ss := (s+s)[1: size * 2  - 1]
	return strings.Contains(ss, s)
}