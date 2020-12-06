package easy

func Reverse(x int) int {
	var ans = 0
	var isLessThanZero = false
	if x < 0 {
		x = -x
		isLessThanZero = true
	}
	for x > 0 {
		ans = ans * 10 + x % 10
		if ans > (2 << 30) {
			return 0
		}
		x = x / 10
	}
	if isLessThanZero {
		ans = - ans
	}
	if ans > (2 << 30) || ans <  - (2 << 30) {
		return 0
	}
	return ans
}
