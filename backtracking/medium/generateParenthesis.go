package medium

/*
*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/
func generateParenthesis(n int) []string {
	ans := []string{}
	dfsParenthesis(&ans, "", 0, 0, n)
	return ans
}

func dfsParenthesis(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		cur += "("
		dfsParenthesis(ans, cur, open+1, close, max)
		cur = cur[:len(cur)-1]
	}
	if close < open {
		cur += ")"
		dfsParenthesis(ans, cur, open, close+1, max)
		cur = cur[:len(cur)-1]
	}
}
