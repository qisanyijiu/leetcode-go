package medium

/*
*
无重复字符的最长子串

https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func LengthOfLongestSubstring(s string) int {
	lastIdx := make([]int, 128)
	// 初始化为 -1
	for i := range lastIdx {
		lastIdx[i] = -1
	}

	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		if lastIdx[c] >= left {
			left = lastIdx[c] + 1
		}
		lastIdx[c] = right
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
