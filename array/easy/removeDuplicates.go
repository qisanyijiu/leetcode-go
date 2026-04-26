package easy

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 { //特殊
		return n
	}
	var slow = 0
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[slow] {
			slow += 1
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
