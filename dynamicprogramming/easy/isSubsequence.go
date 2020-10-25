package easy

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例 1:
s = "abc", t = "ahbgdc"

返回 true.

示例 2:
s = "axc", t = "ahbgdc"

返回 false.
 */
func IsSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	f := make([][26]int, m + 1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j + 'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i + 1][j]
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i] - 'a')] == m {
			return false
		}
		add = f[add][int(s[i] - 'a')] + 1
	}
	return true
}