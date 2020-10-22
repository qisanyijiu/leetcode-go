package easy

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	cases := []struct{
		input, output string
	}{
		struct{ input, output string }{input: "(()())(())", output: "()()()"},
	}
	for _, item := range cases{
		out := RemoveOuterParentheses(item.input)
		if out != item.output{
			t.Errorf("Input: %s, Except: %s, Output: %s\n", item.input, item.output, out)
		}
	}
}