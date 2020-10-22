package medium

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	maxarea, l, r := 0, 0, len(height) - 1
	for l < r {
		maxarea = max(maxarea,  min(height[r], height[l]) * (r - l))
		if height[l] < height[r] {
			l ++
		}else{
			r --
		}
	}
	return maxarea
}