package easy

/**
剑指 Offer 59 - I. 滑动窗口的最大值
https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 */
func MaxSlidingWindow(nums []int, k int) []int {
	var result []int
	var max int

	for i, j := 0, k-1; j >= 0 && j < len(nums); j++ {
		if i == 0 || max == nums[i-1] {
			max = nums[i]
			for t := j; t > i; t-- {
				if max < nums[t] {
					max = nums[t]
				}
			}
		} else {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
		i++
	}
	return result
}