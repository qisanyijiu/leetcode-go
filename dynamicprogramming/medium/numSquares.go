package medium

import "math"

func NumSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt64
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}
