package medium

import "testing"

func Test_NextPermutation(t *testing.T) {
	type TestCase struct {
		Input    []int
		Expected []int
	}
	testCases := []TestCase{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}
	for _, tc := range testCases {
		NextPermutation(tc.Input)
		if !Equal(tc.Expected, tc.Input) {
			t.Errorf("NextPermutation(%v), want %v", tc.Input, tc.Expected)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
