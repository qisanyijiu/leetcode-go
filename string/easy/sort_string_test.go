package easy

import "testing"

func TestSortString(t *testing.T) {
	type Item struct {
		Input string
		Except string
	}
	
	cases := []Item{
		{
			Input:  "aaaabbbbcccc",
			Except: "abccbaabccba",
		},
		{
			Input:  "rat",
			Except: "art",
		},
		{
			Input:  "leetcode",
			Except: "cdelotee",
		},
		{
			Input:  "ggggggg",
			Except: "ggggggg",
		},
		{
			Input:  "spo",
			Except: "ops",
		},
	}
	for _, item := range cases {
		out := SortString(item.Input)
		if out != item.Except {
			t.Errorf("input: %s, except: %s, got: %s", item.Input, item.Except, out)
		}
	}
}