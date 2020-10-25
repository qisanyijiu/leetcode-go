package easy

import "leetcode/utils"

func longestCommonSubstring(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	if lenA == 0 || lenB == 0{
		return 0
	}
	aArr, bArr := []byte(a), []byte(b)
	dp := make([][]int, lenB)
	for i := 0 ; i <lenB; i ++ {
		dp[i] = make([]int, lenA)
	}
	max := 0
	for i := 0; i < lenA; i ++ {
		if bArr[0] == aArr[i]{
			dp[0][i] = 1
			max = 1
		}
	}
	for i := 1; i < lenB; i ++ {
		for j := 0; j < lenA; j ++ {
			if bArr[i] == aArr[j] {
				if j > 0{
					dp[i][j] = dp[i-1][j-1] + 1
				}else{
					dp[i][0] = 1
				}
			}
			max = utils.MaxInt(dp[i][j], max)
		}
	}
	return max
}
