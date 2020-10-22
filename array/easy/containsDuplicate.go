package easy

func ContainsDuplicate(nums []int) bool{
	cache := map[int]int{}
	for _, item := range nums{
		if _, ok := cache[item];ok{
			return true
		}else{
			cache[item] = 1
		}
	}
	return false
}
