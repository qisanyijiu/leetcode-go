package medium

import (
	"leetcode/utils"
)

func MaxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var ans = 0
	for left < right {
		tmp := (right - left) * utils.MinInt(height[left], height[right])
		if tmp > ans {
			ans = tmp
		}
		if height[left] < height[right] {
			left ++
		}else{
			right --
		}
	}
	return ans
}