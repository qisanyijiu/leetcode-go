package easy

func Rotate(nums []int, k int) {
	k %= len(nums)

	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}
func reverse(nums []int, start, end int) {
	for (start < end) {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/**
n = 7
k = 3
arr = [1 2 3 4 5 6 7]

step-1 :
arr = [4 3 2 1 5 6 7]

step-2 :
arr = [4 3 2 1 7 6 5]

step-3 :
arr = [5 6 7 1 2 3 4]
 */