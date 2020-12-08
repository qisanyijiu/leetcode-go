package easy

import "testing"

func TestTitleToNumber(t *testing.T) {
	type Item struct {
		Input string
		Except int
	}
	cases := []Item{
		{
			Input:  "A",
			Except: 1,
		},
		{
			Input:  "AB",
			Except: 28,
		},
		{
			Input:  "ZY",
			Except: 701,
		},
	}
	for _, item := range cases {
		out := TitleToNumber(item.Input)
		if out != item.Except {
			t.Errorf("input: %s, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}