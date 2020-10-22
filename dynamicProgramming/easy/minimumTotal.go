package easy

import (
	"leetcode/utils"
	"math"
)



/**
核心思想如下：从第二行开始逐列遍历，寻找到达该行该列的最短路径。
最后找出最后一行的最小值就可以了
例子：
三角形如下
[2],
[3,4],
[6,5,7],
[4,1,8,3]
路径转移矩阵如下
[2],
[5,6],
[11,10,13],
[14,11,18,16]
最后一行的最小值11就是最短路径
 */
func minimumTotal(triangle [][]int) int{
	depth := len(triangle)
	if depth == 0{
		return 0
	}
	// 初始化dp矩阵
	dp := make([][]int, depth)
	for i := 0; i < depth; i ++ {
		dp[i] = make([]int, depth)
	}
	dp[0][0] = triangle[0][0]
	// 逐行遍历
	for i := 1; i < depth; i ++ {
		// 逐列遍历
		for j := 0; j < len(triangle[i]); j ++ {
			// 防止越界
			lower := utils.MaxInt(0, j - 1)
			upper := utils.MinInt(j , len(triangle[i-1]) - 1)
			// 状态转移
			dp[i][j] = utils.MinInt(dp[i-1][lower], dp[i-1][upper]) + triangle[i][j]
		}
	}
	minCost := int(math.MaxInt32)
	for i := 0; i < len(triangle[depth-1]); i ++ {
		minCost = utils.MinInt(minCost, dp[depth-1][i])
	}
	return minCost
}
