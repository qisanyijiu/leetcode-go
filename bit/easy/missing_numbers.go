package easy

/**
异或运算
 */
func MissingNumber(nums []int) int {
	missing := len(nums)
	for i := range nums {
		missing ^= i ^ nums[i]
	}
	return missing
}