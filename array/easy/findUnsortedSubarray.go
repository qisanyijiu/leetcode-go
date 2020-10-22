package easy

func FindUnsortedSubarray(nums []int) int {
	l := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			l = i-1
			break
		}
	}
	if l == len(nums) {
		return 0
	}
	r := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			r = i + 1
			break
		}
	}
	min, max := nums[l], nums[r]
	for i := l; i <= r; i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
	}
	l, r = -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > min && l == -1 {
			l = i
		}
		if nums[i] < max {
			r = i
		}
	}
	return r - l + 1
}