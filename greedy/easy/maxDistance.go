package easy

/**
https://leetcode.cn/problems/two-furthest-houses-with-different-colors/?envType=daily-question&envId=2026-04-26

街上有 n 栋房子整齐地排成一列，每栋房子都粉刷上了漂亮的颜色。给你一个下标从 0 开始且长度为 n 的整数数组 colors ，其中 colors[i] 表示第  i 栋房子的颜色。

返回 两栋 颜色 不同 房子之间的 最大 距离。

第 i 栋房子和第 j 栋房子之间的距离是 abs(i - j) ，其中 abs(x) 是 x 的绝对值。
*/
func maxDistance(colors []int) int {
	n := len(colors)
	maxDist := 0
	max := func(a, b, c int) int {
		res := a
		if res < b {
			res = b
		}
		if res < c {
			res = c
		}
		return res
	}
	for i := 1; i < n; i ++ {
		if colors[i] != colors[0] || colors[i] != colors[n-1] {
			maxDist = max(maxDist, i, n-1-i)
		}
	}
	return maxDist
}