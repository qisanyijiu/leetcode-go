package medium

import (
	"testing"
)

func Test_FindDuplicate(t *testing.T) {
	type TestCase struct {
		input  []int
		Except int
	}
	cases := []TestCase{
		{input: []int{1, 2, 3, 2}, Except: 2},
		{input: []int{1, 2, 3, 3}, Except: 3},
		{input: []int{1, 1, 1, 1, 1}, Except: 1},
	}
	for _, c := range cases {
		output := FindDuplicate(c.input)
		if output != c.Except {
			t.Errorf("Find Duplicate output: %v, expect: %v", output, c.Except)
		}
	}
}
