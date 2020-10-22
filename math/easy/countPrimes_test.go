package easy

import "testing"

func TestCountPrimes(t *testing.T) {
	cases := []struct{
		input int
		output int
	}{
		{input: 10, output: 4},
		{input: 0, output: 0},
		{input: 1, output: 0},
		{input: 2, output: 0},
	}
	for _, item := range cases{
		out := CountPrimes(item.input)
		if out != item.output {
			t.Errorf("input: %v, except: %v, got: %v", item.input, item.output, out)
		}
	}

}