package easy

/**
你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。

给定一个数字 n，找出可形成完整阶梯行的总行数。

n 是一个非负整数，并且在32位有符号整型的范围内。

示例 1:

n = 5

硬币可排列成以下几行:
¤
¤ ¤
¤ ¤

因为第三行不完整，所以返回2.
示例 2:

n = 8

硬币可排列成以下几行:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

因为第四行不完整，所以返回3.
*/
func ArrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	var (
		ans = 1
		min = 1
		max = n
	)
	for min <= max {
		var mid = (min + max) / 2
		var sum = mid * (mid + 1)
		if sum == 2 * n {
			return mid
		}else if sum < 2 * n {
			ans = mid
			min = mid + 1
		}else{
			max = mid - 1
		}
	}
	return ans
}
