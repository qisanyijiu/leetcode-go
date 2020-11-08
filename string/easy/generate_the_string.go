package easy

import "strings"

/**
1374. 生成每种字符都是奇数个的字符串

https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts/

给你一个整数 n，请你返回一个含 n 个字符的字符串，其中每种字符在该字符串中都恰好出现 奇数次 。

返回的字符串必须只含小写英文字母。如果存在多个满足题目要求的字符串，则返回其中任意一个即可。

 

示例 1：

输入：n = 4
输出："pppz"
解释："pppz" 是一个满足题目要求的字符串，因为 'p' 出现 3 次，且 'z' 出现 1 次。当然，还有很多其他字符串也满足题目要求，比如："ohhh" 和 "love"。
示例 2：

输入：n = 2
输出："xy"
解释："xy" 是一个满足题目要求的字符串，因为 'x' 和 'y' 各出现 1 次。当然，还有很多其他字符串也满足题目要求，比如："ag" 和 "ur"。
示例 3：

输入：n = 7
输出："holasss"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-a-string-with-characters-that-have-odd-counts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func generateTheString(n int) string {
	var str strings.Builder
	c :=1
	if n>1&&n%2!=0{
		c=2
	}
	char :=[]byte{'a','b','c'}
	for i:=0;i<c;i++{
		str.WriteByte(char[i])
	}
	for i:=0;i<n-c;i++{
		str.WriteByte(char[c])
	}
	return str.String()
}