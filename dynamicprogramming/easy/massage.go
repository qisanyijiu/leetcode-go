package easy

func Massage(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2{
		return max(nums[1], nums[0])
	}
	nums[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i ++ {
		nums[i] = max(nums[i-1], nums[i] + nums[i-2])
	}
	return nums[len(nums) - 1]
}