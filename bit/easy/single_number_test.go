package easy

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{input: []int{2, 2, 1}, output: 1},
		{input: []int{4, 1, 2, 1, 2}, output: 4},
	}
	for _, item := range cases {
		res := SingleNumber(item.input)
		if res != item.output {
			t.Errorf("input: %v, except: %d, got: %d", item.input, item.output, res)
		}
	}
}
