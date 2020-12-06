package easy

import "testing"

func TestReverse(t *testing.T) {
	type Item struct {
		Input int
		Except int
	}
	cases := []Item{
		{
			Input:  123,
			Except: 321,
		},
		{
			Input:  -123,
			Except: -321,
		},
		{
			Input:  120,
			Except: 21,
		},
		{
			Input: 1563847412,
			Except: 0,
		},
	}
	for _, item := range cases{
		out := Reverse(item.Input)
		if out != item.Except {
			t.Errorf("input: %d, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}