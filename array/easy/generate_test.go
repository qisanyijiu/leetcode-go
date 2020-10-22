package easy

import "testing"

func TestGenerate(t *testing.T) {
	cases := []struct{
		input int
		output [][]int
	}{
		struct {
			input  int
			output [][]int
		}{input: 1, output: [][]int{[]int{1}}},
		struct {
			input  int
			output [][]int
		}{input: 2, output: [][]int{[]int{1}, []int{1, 1}}},
		struct {
			input  int
			output [][]int
		}{input: 3, output: [][]int{[]int{1}, []int{1, 2, 1}}},
		struct {
			input  int
			output [][]int
		}{input: 5, output: [][]int{[]int{1}, []int{1, 4, 6, 4, 1}}},
	}
	for _, item := range cases{
		out := Generate(item.input)
		if !Arr2Equal(out, item.output){
			t.Errorf("Input: %d, Output: %v, Except: %v;", item.input, out,item.output)
		}
	}
}

func Arr2Equal(arr1, arr2 [][]int) bool {
	if len(arr1) != len(arr2){
		return false
	}
	for rowIndex, row1 := range arr1{
		row2 := arr2[rowIndex]
		if len(row2) != len(row1){
			return false
		}
		for colIndex, val := range row1{
			if val != row2[colIndex]{
				return false
			}
		}
	}
	return true
}