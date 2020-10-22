package medium

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int{
	abs := func(n int)int{
		if n < 0 {
			return -n
		}else {
			return n
		}
	}
	sort.Ints(nums)
	var sum int= math.MaxInt32
	for i, num := range nums{
		if i > 0 && num == nums[i-1]{
			continue
		}
		l, h := i + 1, len(nums) - 1
		for l < h {
			x := num + nums[l] + nums[h]
			if x == target{
				return x
			}else if x > target{
				h -= 1
			}else {
				l += 1
			}
			if abs(x - target) < abs(sum - target){
				sum = x
			}
		}
	}
	return sum
}
