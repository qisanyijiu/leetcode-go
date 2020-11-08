package easy

import "testing"

func TestCompressString(t *testing.T) {
	type Item struct {
		Input string
		Except string
	}
	cases := []Item{
		{
			Input:  "aabcccccaaa",
			Except: "a2b1c5a3",
		},
		{
			Input:  "abbccd",
			Except: "abbccd",
		},
	}
	for _, item := range cases {
		out := CompressString(item.Input)
		if out != item.Except {
			t.Errorf("input: %s, except: %s, got: %s", item.Input, item.Except, out)
		}
	}
}
