package medium

/**
46. 全排列

https://leetcode-cn.com/problems/permutations/

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var ans = [][]int{}
	backtrace(nums, []int{}, &ans)
	return ans
}

func backtrace(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, prev)
		return
	}

	for i := 0; i < len(nums); i++ {
		backtrace(
			removeElem(nums, i),         // 将数组的第I位取出来
			append(prev, nums[i]), // 将取出来的数字放到前缀中区
			ans)
	}
}

func removeElem(nums []int, index int) []int {
	return append(append([]int{}, nums[0:index]...), nums[index+1:]...)
}
