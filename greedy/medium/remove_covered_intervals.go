package medium

import "sort"

/**
1288. 删除被覆盖区间

https://leetcode-cn.com/problems/remove-covered-intervals/

给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。

只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。

在完成所有删除操作后，请你返回列表中剩余区间的数目。

示例：

输入：intervals = [[1,4],[3,6],[2,8]]
输出：2
解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。
 */
func RemoveCoveredIntervals(intervals [][]int) int {

	// 按照起点升序排列，起点相同时降序排列
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]

	res := 0

	for i :=1; i < len(intervals); i ++ {
		intv := intervals[i]

		// 情况一，找到覆盖区间
		if left <= intv[0] && right >= intv[1] {
			res++;
		}
		// 情况二，找到相交区间，合并
		if right >= intv[0] && right <= intv[1] {
			right = intv[1];
		}
		// 情况三，完全不相交，更新起点和终点
		if right < intv[0] {
			left = intv[0];
			right = intv[1];
		}
	}

	return len(intervals) - res
}
