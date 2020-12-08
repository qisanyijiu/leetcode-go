package easy

import "testing"

func TestIsHappy(t *testing.T) {
	type Item struct {
		Input  int
		Except bool
	}
	cases := []Item{
		{
			Input:  19,
			Except: true,
		},
	}
	for _, item := range cases {
		out := IsHappy(item.Input)
		if out != item.Except {
			t.Errorf("input: %d, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}
