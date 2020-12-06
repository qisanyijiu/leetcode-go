package easy

/**
13. 罗马数字转整数

https://leetcode-cn.com/problems/roman-to-integer/submissions/

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
 */
func RomanToInt(s string) int {
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10。

	// 当右边的字符比左边的大， 说明是特殊组合
	pre, r := 0,0
	for i:=len(s)-1; i>=0;i-- {
		cur := _map[s[i]]
		if cur >= pre {
			r += cur
		}else {
			r -= cur
		}
		pre = cur
	}
	return r
}