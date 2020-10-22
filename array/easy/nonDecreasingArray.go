package easy

func checkPossibility(nums []int) bool {
	index := -1
	length := len(nums)
	for i := 0; i < length - 1; i ++ {
		if nums[i] > nums[i+1]{
			if index > -1{
				return false
			}
			index = i
		}
	}
	return index == - 1 || index == 0 || index == len(nums) - 2 || nums[index-1] <= nums[index+1] || nums[index] <= nums[index+2]
}