package medium

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minn := math.MaxInt64
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] >= 0 {
				minn = min(minn, dp[i-coin]+1)
			}
		}
		if minn == math.MaxInt64 {
			dp[i] = -1
		} else {
			dp[i] = minn
		}
	}
	return dp[amount]
}
