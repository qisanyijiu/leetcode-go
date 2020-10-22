package easy

func RemoveElement(nums []int, val int) int {
	cnt := len(nums)
	for i := 0; i < cnt; i ++ {
		item := nums[i]
		if item == val{
			nums = append(nums[:i], nums[i+1:]...)
			i --
			cnt --
		}
	}
	return cnt
}
