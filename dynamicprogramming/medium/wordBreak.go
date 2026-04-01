package medium

func wordBreak(s string, wordDict []string) bool {
	wMap := make(map[string]bool)
	for _, w := range wordDict {
		wMap[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
