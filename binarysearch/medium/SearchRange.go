package medium

func SearchRange(nums []int, target int) []int {
	leftBound := searchLeftBound(nums, target)
	rightBound := searchRightBound(nums, target)
	return []int{leftBound, rightBound}
}

func searchLeftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			right  = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if (left >= len(nums) || nums[left] != target) {
		return -1
	}
	return left
}

func searchRightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return left - 1
}
