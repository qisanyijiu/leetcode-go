package medium

import (
	"math"
	"sort"
)

/**
16. 最接近的三数之和

https://leetcode-cn.com/problems/3sum-closest/

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 */
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n = len(nums)
		best = math.MaxInt32
	)
	update := func(cur int) {
		if abs(cur - target) < abs(best - target) {
			best = cur
		}
	}
	for i := 0; i < n ; i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i + 1, n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0 --
				}
				k = k0
			}else{
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0 ++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}