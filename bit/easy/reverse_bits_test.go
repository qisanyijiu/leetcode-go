package easy

import "testing"

func TestReverseBits(t *testing.T) {
	type Item struct {
		Input uint32
		Except uint32
	}
	cases := []Item{
		{
			Input:  43261596,
			Except: 964176192,
		},
		{
			Input:  4294967293,
			Except: 3221225471,
		},
	}
	for _, item := range cases {
		out := ReverseBits(item.Input)
		if out != item.Except {
			t.Errorf("input: %d, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}
