package medium

/**
986. 区间列表的交集

https://leetcode-cn.com/problems/interval-list-intersections/

给定两个由一些 闭区间 组成的列表，每个区间列表都是成对不相交的，并且已经排序。

返回这两个区间列表的交集。

（形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b。两个闭区间的交集是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3]。）
 */

func IntervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	res := [][]int{}
	for i < len(A) && j < len(B){
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]

		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if b2 < a2 {
			j ++
		}else {
			i ++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

