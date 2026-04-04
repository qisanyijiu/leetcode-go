package medium

import "math"

func MaxProfit(prices []int) int {
	ans, preMin := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < preMin {
			preMin = prices[i]
		} else {
			if prices[i]-preMin > ans {
				ans = prices[i] - preMin
			}
		}
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
