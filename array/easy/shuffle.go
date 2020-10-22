package easy

/**
打乱一个没有重复元素的数组。

 

示例:

// 以数字集合 1, 2 和 3 初始化数组。
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
solution.shuffle();

// 重设数组到它的初始状态[1,2,3]。
solution.reset();

// 随机返回数组[1,2,3]打乱后的结果。
solution.shuffle();
 */

import "math/rand"

type Solution struct {
	nums []int
	original []int
}


func Constructor(nums []int) Solution {
	var s Solution
	s.nums=nums
	s.original=make([]int, len(nums))
	copy(s.original, nums)
	return s
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	len:=len(this.nums)
	for index,_:=range this.nums{
		swapIndex := random(index, len)
		this.nums[index], this.nums[swapIndex] = this.nums[swapIndex], this.nums[index]
	}
	return this.nums
}

// [min, max)
func random(min, max int) int{
	return rand.Intn(max-min)+min
}
