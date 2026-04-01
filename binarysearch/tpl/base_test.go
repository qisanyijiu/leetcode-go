package tpl

import "testing"

func TestBinarySearch(t *testing.T) {
	type Case struct {
		Data     []int
		Target   int
		Expected int
	}
	cases := []Case{
		{
			Data:     []int{1, 2, 3, 4, 5},
			Target:   1,
			Expected: 0,
		},
		{
			Data:     []int{1, 2, 3, 4, 5},
			Target:   5,
			Expected: 4,
		},
		{
			Data:     []int{1, 2, 3, 4, 5},
			Target:   3,
			Expected: 2,
		},
		{
			Data:     []int{1, 2, 3, 4, 5},
			Target:   0,
			Expected: -1,
		},
		{
			Data:     []int{1, 2, 3, 4, 5},
			Target:   6,
			Expected: -1,
		},
	}
	for _, c := range cases {
		if got := BinarySearch(c.Data, c.Target); got != c.Expected {
			t.Errorf("BinarySearch(%v, %v): got %v, want %v", c.Data, c.Target, got, c.Expected)
		}
	}
}
