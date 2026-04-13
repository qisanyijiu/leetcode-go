package medium

/*
*
https://leetcode.cn/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/description/?envType=daily-question&envId=2026-04-11
给你一个 m x n 的网格图，其中 (0, 0) 是最左上角的格子，(m - 1, n - 1) 是最右下角的格子。给你一个整数数组 startPos ，startPos = [startrow, startcol] 表示 初始 有一个 机器人 在格子 (startrow, startcol) 处。同时给你一个整数数组 homePos ，homePos = [homerow, homecol] 表示机器人的 家 在格子 (homerow, homecol) 处。

机器人需要回家。每一步它可以往四个方向移动：上，下，左，右，同时机器人不能移出边界。每一步移动都有一定代价。再给你两个下标从 0 开始的整数数组：长度为 m 的数组 rowCosts  和长度为 n 的数组 colCosts 。

如果机器人往 上 或者往 下 移动到第 r 行 的格子，那么代价为 rowCosts[r] 。
如果机器人往 左 或者往 右 移动到第 c 列 的格子，那么代价为 colCosts[c] 。
请你返回机器人回家需要的 最小总代价 。
*/
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	r1, c1 := startPos[0], startPos[1]
	r2, c2 := homePos[0], startPos[1]
	res := 0
	if r2 >= r1 {
		for i := r1 + 1; i <= r2; i++ {
			res += rowCosts[i]
		}
	} else {
		for i := r2; i < r1; i++ {
			res += rowCosts[i]
		}
	}
	if c2 >= c1 {
		for i := c1 + 1; i <= c2; i++ {
			res += colCosts[i]
		}
	} else {
		for i := c2; i < c1; i++ {
			res += colCosts[i]
		}
	}
	return res
}
