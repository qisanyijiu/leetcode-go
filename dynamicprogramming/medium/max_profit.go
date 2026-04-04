package medium

import "math"

func MaxProfit(prices []int) int {
	var ans = 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

func max(arr ...int) int {
	cur := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if cur < arr[i] {
			cur = arr[i]
		}
	}
	return cur
}
