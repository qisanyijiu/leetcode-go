package easy


func Generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i ++ {
		arr := make([]int, i + 1)
		for j := 0; j < i+1; j++{
			// 每行第一个和最后一个为1
			if j == 0 || j == i{
				arr[j] = 1
			}else {
				arr[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = arr
	}
	return result
}