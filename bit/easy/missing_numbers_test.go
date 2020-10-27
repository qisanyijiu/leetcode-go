package easy

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{input: []int{0, 2, 1}, output: 3},
		{input: []int{0,1}, output: 2},
	}
	for _, item := range cases {
		res := MissingNumber(item.input)
		if res != item.output {
			t.Errorf("input: %v, except: %d, got: %d", item.input, item.output, res)
		}
	}
}