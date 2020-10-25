package easy

import "testing"

func TestDivisorGame(t *testing.T) {
	cases := []struct {
		Input int
		Out   bool
	}{
		{Input: 2, Out: true},
		{Input: 3, Out: false},
		{Input: 2, Out: true},
		{Input: 4, Out: true},
		{Input: 5, Out: false},
	}
	for _, item := range cases {
		res := DivisorGame(item.Input)
		if res != item.Out {
			t.Errorf("input: %d, except: %t, got :%t", item.Input, item.Out, res)
		}
	}
}
