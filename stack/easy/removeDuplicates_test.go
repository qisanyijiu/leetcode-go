package easy

import "testing"

func TestRemoveDuplicates(t *testing.T){
	cases := []struct{
		input, output string
	}{
		struct{ input, output string }{input: "abbaca", output: "ca"},
		struct{ input, output string }{input: "fuck", output: "fuck"},
	}
	for _, item := range cases{
		out := RemoveDuplicates(item.input)
		if out != item.output{
			t.Errorf("Input: %s, Except: %s, Output: %s\n", item.input, item.output, out)
		}
	}
}
