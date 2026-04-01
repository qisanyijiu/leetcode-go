package easy

/*
*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

测试用例保证输出的长度不会超过 105。

输入：s = "3[a]2[bc]"
输出："aaabcbc"
*/
func decodeString(s string) string {
	stack := []byte{}
	for _, b := range []byte(s) {
		if b == ']' {
			j := len(stack) - 1
			for j = len(stack) - 1; j >= 0; j-- {
				if stack[j] == '[' {
					break
				}
			}
			chars := stack[j+1:]
			cnt := int(stack[j-1] - '0')
			buf := make([]byte, 0)
			for i := 0; i < cnt; i++ {
				buf = append(buf, chars...)
			}
			stack = append(stack, buf...)
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}
