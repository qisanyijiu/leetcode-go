package hard

/**
串联所有单词的子串

给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

todo: 没有理解
*/
func FindSubstring(s string, words []string) []int {
	var res = []int{}
	if len(words) == 0 {
		return res
	}

	var wLen = len(words[0])
	var sLen = len(s)
	// 确定字串的长度
	var tLen = wLen * len(words)
	if sLen < tLen {
		return res
	}

	var wordDict = make(map[string]int)
	for _, word := range words {
		wordDict[word] ++
	}

	// 移动窗口减少重复检查单词，按单词长度取不同批次
	for i := 0; i < wLen; i++ {
		var count int
		var countermp = make(map[string]int)

		for l, r := i, i; r <= sLen - wLen; r = r + wLen{
			word := s[r:r + wLen]
			if num, ok := wordDict[word]; ok {
				// 如果计数器中单词数目超标，左移指针直至符合数目要求
				for countermp[word] >= num {
					countermp[s[l:l+wLen]]--
					count--
					l+=wLen
				}
				countermp[word]++
				count++
			}else{
				// 如果当前单词不在词典里，左移指针至下一个单词，左移过程中清理计数
				for l < r {
					countermp[s[l:l+wLen]] --
					count --
					l += wLen
				}
				l += wLen
			}
			if count == len(words) {
				res = append(res, l)
			}
		}
	}
	return res
}
