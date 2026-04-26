package easy

func getMinDistance(nums []int, target int, start int) int {
	left := findLeft(nums, target, start)
	right := findRight(nums, target, start)
	minDist := len(nums)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if left != -1 {
		minDist = min(minDist, start - left)
	}
	if right != -1 {
		minDist = min(minDist, right - start)
	}
	return minDist
}

func findLeft(nums []int, target int, start int) int {
	for i := start; i >= 0; i -- {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func findRight(nums []int, target int, start int) int {
	for i := start; i < len(nums); i ++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}