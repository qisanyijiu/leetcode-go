package easy

/**
爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。

最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：

选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。

只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 False。假设两个玩家都以最佳状态参与游戏。
 */

func DivisorGame(N int) bool {
	return divisorGameBit(N)
}

func divisorGameDp(N int) bool {
	dp := make([]bool, N+2)
	dp[2] = true
	for i := 3; i < N+1; i ++ {
		for j := 1; j < i; j ++ {
			if i % j == 0 && dp[i-j] == false{
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

func divisorGameBit(N int) bool {
	return (N & 1) != 1
}