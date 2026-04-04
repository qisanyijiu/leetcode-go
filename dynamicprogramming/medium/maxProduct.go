package medium

/*
*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

请注意，一个只包含一个元素的数组的乘积是这个元素的值。
*/
func maxProduct(nums []int) int {
	preMax, preMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := preMax, preMin
		preMax = max(nums[i], mx*nums[i], mn*nums[i])
		preMin = min(nums[i], mn*nums[i], mx*nums[i])
		if preMax > ans {
			ans = preMax
		}
	}
	return ans
}
