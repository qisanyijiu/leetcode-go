package medium

/*
*

https://leetcode.cn/problems/unique-paths/description/?envType=study-plan-v2&envId=top-100-liked

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
*/
func UniquePaths(row, col int) int {
	// 只有一行或者一列时，只有一种路径
	if row == 1 || col == 1 {
		return 1
	}
	// 只记录上一行就行
	dp := make([]int, col)
	// 第一列全为1
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[col-1]
}
