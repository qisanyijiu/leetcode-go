package easy

import "math"

/**
69. x 的平方根

https://leetcode-cn.com/problems/sqrtx/

实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
 */
func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	var (
		min = 1
		max = x / 2
		ans = -1
	)
	for min <= max {
		var mid = min + (max  - min) / 2
		if mid * mid <= x {
			ans = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return ans
}

/**
牛顿迭代
 */
func NewtonSqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	C , x0 := float64(x), float64(x) / 2
	for {
		xi := 0.5 * (x0 + C /x0)
		if math.Abs(x0 - xi) < 0.09 {
			break
		}
		x0 = xi
	}
	return int(x0)
}