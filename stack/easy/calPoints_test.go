package easy

import "testing"

func TestCalPoints(t *testing.T) {
	cases := []struct {
		input []string
		output int
	}{
		struct {
			input  []string
			output int
		}{input: []string{"5","2","C","D","+"}, output: 30},
		struct {
			input  []string
			output int
		}{input: []string{"5","-2","4","C","D","9","+","+"}, output: 27},
	}
	for _, item := range cases{
		result := CalPoints(item.input)
		if result != item.output{
			t.Errorf("Input: %v, Except: %d, Got: %d\n", item.input, item.output, result)
		}
	}
}