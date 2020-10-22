package easy

func MissingNumber(nums []int) int {
	var counter, shouldSum, sum int

	for _, v := range nums {
		sum += v
		shouldSum += counter
		counter++
	}

	shouldSum += counter

	return shouldSum - sum
}