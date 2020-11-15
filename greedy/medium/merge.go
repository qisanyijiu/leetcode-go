package medium

import "sort"

/**
56. 合并区间

https://leetcode-cn.com/problems/merge-intervals/

给出一个区间的集合，请合并所有重叠的区间。
示例 1:

输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 */
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res = [][]int{}
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i ++ {
		curr := intervals[i]
		last := res[len(res)-1]
		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
			res[len(res)-1] = last
		}else{
			res = append(res, curr)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}