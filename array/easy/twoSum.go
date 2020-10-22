package easy

func TwoSum(nums []int, target int) []int {
	dict := map[int]int{}
	result := make([]int, 2)
	for i, val := range nums{
		condition := target - val
		if index, ok := dict[condition]; ok {
			result = []int{index, i}
			break
		}else{
			dict[val] = i
		}
	}
	return result
}