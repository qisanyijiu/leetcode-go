package medium

/**
https://leetcode.cn/problems/partition-equal-subset-sum/description/?envType=study-plan-v2&envId=top-100-liked

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

示例 1：
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

示例 2：
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
*/

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum, maxNum := 0, 0
	for _, n := range nums {
		sum += n
		maxNum = max(maxNum, sum)
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 1; j <= target; j++ {
			if j >= num {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
