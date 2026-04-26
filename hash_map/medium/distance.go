package medium

/**
https://leetcode.cn/problems/sum-of-distances/?envType=daily-question&envId=2026-04-26

给你一个下标从 0 开始的整数数组 nums 。现有一个长度等于 nums.length 的数组 arr 。对于满足 nums[j] == nums[i] 且 j != i 的所有 j ，arr[i] 等于所有 |i - j| 之和。如果不存在这样的 j ，则令 arr[i] 等于 0 。

返回数组 arr 。

 

示例 1：

输入：nums = [1,3,1,1,2]
输出：[5,0,3,4,0]
解释：
i = 0 ，nums[0] == nums[2] 且 nums[0] == nums[3] 。因此，arr[0] = |0 - 2| + |0 - 3| = 5 。 
i = 1 ，arr[1] = 0 因为不存在值等于 3 的其他下标。
i = 2 ，nums[2] == nums[0] 且 nums[2] == nums[3] 。因此，arr[2] = |2 - 0| + |2 - 3| = 3 。
i = 3 ，nums[3] == nums[0] 且 nums[3] == nums[2] 。因此，arr[3] = |3 - 0| + |3 - 2| = 4 。 
i = 4 ，arr[4] = 0 因为不存在值等于 2 的其他下标。
示例 2：

输入：nums = [0,5,3]
输出：[0,0,0]
解释：因为 nums 中的元素互不相同，对于所有 i ，都有 arr[i] = 0 。
​
 */
func distance(nums []int) []int64 {
    n := len(nums)
	groups := make(map[int][]int, 0)
	for idx, num := range nums {
		if _, ok := groups[num]; ok {
			groups[num] = append(groups[num], idx)
		}else{
			groups[num] = []int{idx}
		}
	}
	res := make([]int64, n)
	for _, group := range groups {
		var total int64 = 0
		for _, idx := range group {
			total += int64(idx - group[0])
		}
		var prefixTotal int64
		for i, idx := range group {
			res[idx] = total - prefixTotal * 2 + int64(idx) * int64(2 * i - len(group))
            prefixTotal += int64(idx)
		}
	}
	return res
}