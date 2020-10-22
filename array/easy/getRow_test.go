package easy

import "testing"

func TestGetRow(t *testing.T) {
	cases := []struct{
		input int
		output []int
	}{
		struct {
			input  int
			output []int
		}{input: 1, output: []int{1,1}},
		struct {
			input  int
			output []int
		}{input: 2, output: []int{1, 2, 1}},
		struct {
			input  int
			output []int
		}{input: 3, output: []int{1, 3, 3, 1}},
		struct {
			input  int
			output []int
		}{input: 4, output:  []int{1, 4, 6, 4, 1}},
	}
	for _, item := range cases{
		out := GetRow(item.input)
		if !ArrEqual(out, item.output){
			t.Errorf("Input: %d, Output: %v, Except: %v;", item.input, out,item.output)
		}
	}
}