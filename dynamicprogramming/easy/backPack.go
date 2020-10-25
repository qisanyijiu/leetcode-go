package easy

import (
	"leetcode/utils"
)

/**
核心思路，假设背包的容量为o

 */
func backPack(m int, A []int) int {
	length := len(A)
	if length == 0 || m == 0 {
		return 0
	}

	dp := make([][]int , length)
	for i := 0; i < length; i ++ {
		dp[i] = make([]int, m+1)
	}
	for j := 0; j < m + 1;  j ++{
		if j >= A[0]{
			dp[0][j] = A[0]
		}
	}
	// 尝试将物品塞入
	for i := 1; i < length; i ++ {
		// 增加包的容量
		for j := 1; j < m + 1; j ++ {
			if j >= A[i] {
				dp[i][j] = utils.MaxInt(dp[i-1][j], dp[i-1][j-A[i]] + A[i])
			}else{
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[length-1][m]
}