package medium

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	start := 0
	end := len(nums) - 1

	for i := 0; i <= end; {
		if nums[i] == 0 {
			if i != start {
				nums[start], nums[i] = nums[i], nums[start]
			}
			start++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			i++
		}
	}
}
