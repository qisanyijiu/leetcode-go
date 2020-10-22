package easy

import "testing"

func TestTwoSum2(t *testing.T) {
	cases := []struct{
		input, output []int
		target int
	}{
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{1,2}, target: 9},
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{1,3}, target: 13},
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{2, 3}, target: 18},
	}
	for _, item := range cases{
		result := TwoSum2(item.input, item.target)
		if !twoSumEqual(result, item.output){
			t.Errorf("Array: %v, Target: %d, Except: %v, Got: %v\n", item.input, item.target, item.output, result)
		}
	}
}