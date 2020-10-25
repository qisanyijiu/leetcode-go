package easy

import "testing"

func TestMassage(t *testing.T) {
	cases := []struct {
		Input []int
		Out   int
	}{
		{Input: []int{1, 2, 3, 1}, Out: 4},
		{Input: []int{2, 7, 9, 3, 1}, Out: 12},
		{Input: []int{2, 1, 4, 5, 3, 1, 1, 3}, Out: 12},
	}
	for _, item := range cases {
		res := Massage(item.Input)
		if res != item.Out {
			t.Errorf("input: %v, except: %d, got: %d", item.Input, item.Out, res)
		}
	}
}
