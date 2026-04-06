package medium

/*
*
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
	trace := []int{}
	used := make([]bool, len(nums))
	backtrace(nums, trace, used, &ans)
	return ans
}

func backtrace(nums []int, trace []int, used []bool, ans *[][]int) {
	if len(nums) == len(trace) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*ans = append(*ans, temp)
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			// 数字已经使用了，跳过
			continue
		}
		// 没有使用过，数字加入路径
		trace = append(trace, nums[i])
		// 标记数字已经使用过了
		used[i] = true
		// 进入下一层决策树
		backtrace(nums, trace, used, ans)
		// 路径探索完后进行回退
		trace = trace[:len(trace)-1]
		// 标记为为使用q
		used[i] = false
	}
}
