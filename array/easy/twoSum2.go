package easy

func TwoSum2(numbers []int, target int) []int {
	dict := map[int]int{}
	result := make([]int, 2)
	for i, val := range numbers{
		condition := target - val
		if index, ok := dict[condition]; ok {
			result = []int{index+1, i+1}
			break
		}else{
			dict[val] = i
		}
	}
	return result
}