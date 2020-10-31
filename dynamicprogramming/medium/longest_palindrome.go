package medium

/**
5.最长回文子串
https://leetcode-cn.com/problems/longest-palindromic-substring/

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
func LongestPalindrome(s string) string {
	return LongestPalindromeDP(s)
}

// 动态规划算法
func LongestPalindromeDP(s string) string {
	var length = len(s)
	var res = ""
	//初始化dp
	var dp = make([][]int, length)
	for i := 0; i < length; i ++ {
		dp[i] = make([]int, length)
	}
	for l := 0; l < length; l++ {
		for i := 0; i+l < length; i++ {
			j := i + l
			if l == 0 {
				// 如果长度为1，必定为会文串
				dp[i][j] = 1
			} else if l == 1 {
				// 长度为2，如果开始和结束相等，则为回文串
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				// 如果长度大于2，则之前相等，现在也相等
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(res) {
				res = s[i : i+l+1]
			}
		}
	}
	return res
}

// 中心扩展算法
func LongestPalindromeMID(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}

// Manacher 算法
func LongestPalindromeManacher(s string) string {
	return ""
}
