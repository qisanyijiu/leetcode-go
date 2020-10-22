package easy

func GetRow(rowIndex int) []int {
	result := make([]int,0, rowIndex+1)
	result = append(result, 1)
	for i:= 1; i <= rowIndex; i ++ {
		prev := result[0]
		for j := 1; j < len(result); j++ {
			dummy := result[j]
			result[j] = result[j] + prev
			prev = dummy
		}
		result = append(result, 1)
	}
	return result
}