package easy

import "testing"

func TestMySqrt(t *testing.T) {
	type Item struct {
		X int
		Y int
	}
	cases := []Item{
		{
			X: 4,
			Y: 2,
		},
		{
			X: 8,
			Y: 2,
		},
		{
			X: 9,
			Y: 3,
		},
	}

	for _, item := range cases {
		out := MySqrt(item.X)
		if out != item.Y {
			t.Errorf("x: %d, except: %d, got: %d", item.X, item.Y, out)
		}
	}
}
