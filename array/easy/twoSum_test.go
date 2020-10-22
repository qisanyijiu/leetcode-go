package easy

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct{
		input, output []int
		target int
	}{
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{0,1}, target: 9},
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{0,2}, target: 13},
		struct {
			input, output []int
			target        int
		}{input: []int{2, 7, 11, 15}, output: []int{1, 2}, target: 18},
	}
	for _, item := range cases{
		result := TwoSum(item.input, item.target)
		if !twoSumEqual(result, item.output){
			t.Errorf("Array: %v, Target: %d, Except: %v, Got: %v\n", item.input, item.target, item.output, result)
		}
	}
}

func twoSumEqual(nums1, nums2 []int) bool {
	if len(nums1) != 2 || len(nums2) != 2{
		return false
	}
	if nums1[0] != nums2[0] || nums1[1] != nums2[1] {
		return false
	}
	return true
}