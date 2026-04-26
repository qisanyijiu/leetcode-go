package easy

/**
https://leetcode.cn/problems/mirror-distance-of-an-integer/description/?envType=daily-question&envId=2026-04-26

给你一个整数 n。

定义它的 镜像距离 为：abs(n - reverse(n))​​​​​​​，其中 reverse(n) 表示将 n 的数字反转后形成的整数。

返回表示 n 的镜像距离的整数。

其中，abs(x) 表示 x 的绝对值。

 

示例 1：

输入： n = 25

输出： 27

解释：

reverse(25) = 52。
因此，答案为 abs(25 - 52) = 27。
示例 2：

输入： n = 10

输出： 9

解释：
*/
func mirrorDistance(n int) int {
	return abs(n - reverse(n))
}

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = rev * 10 + n % 10
		n /= 10
	}
	return rev
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}