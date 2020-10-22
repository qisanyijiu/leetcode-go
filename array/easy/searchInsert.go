package easy

func SearchInsert(nums []int, target int) int {
	cnt := len(nums)
	for i:= 0; i<cnt; i ++ {
		if target <= nums[i]{
			return i
		}
	}
	return cnt
}
