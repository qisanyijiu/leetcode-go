package easy

func waysToStep(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	dp1, dp2, dp3 := 1, 2, 4
	for i := 4; i <= n; i++ {
		dp1, dp2, dp3 = dp2, dp3, (dp1+dp2+dp3)%1000000007
	}
	return dp3
}