package easy

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	length := len(nums)
	mem := make([]int, length+1)
	for _, v := range nums{
		mem[v] = v
	}
	for i := 1; i < length + 1; i ++ {
		if mem[i] == 0{
			result = append(result, i)
		}
	}
	return result
}