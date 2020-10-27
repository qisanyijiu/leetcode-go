package easy

import "testing"

func TestHammingWeight(t *testing.T) {
	cases := []struct {
		input  uint32
		output int
	}{
		{input: uint32(1), output: 1},
		{input: uint32(4), output: 1},
		{input: uint32(3), output: 2},
	}
	for _, item := range cases {
		res := HammingWeight(item.input)
		if res != item.output {
			t.Errorf("input: %d, except: %d, got: %d", item.input, item.output, res)
		}
	}
}
