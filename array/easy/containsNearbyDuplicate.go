package easy

func ContainsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	flag := map[int]int{}
	for index, value := range nums{
		if _, ok := flag[value]; !ok {
			flag[value] = 0
		}
		if flag[value] > 0 && index - flag[value] + 1 <= k{
			return true
		}
		flag[value] = index + 1
	}
	return false
}