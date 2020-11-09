package easy

import "testing"

func TestArrangeCoins(t *testing.T) {
	type Item struct {
		N     int
		Lines int
	}
	cases := []Item{
		{
			N:     1,
			Lines: 1,
		},
		{
			N:     3,
			Lines: 2,
		},
		{
			N:     5,
			Lines: 2,
		},
		{
			N:     8,
			Lines: 3,
		},
	}
	for _, item := range cases {
		ans := ArrangeCoins(item.N)
		if ans != item.Lines {
			t.Errorf("n: %d, except: %d, got: %d", item.N, item.Lines, ans)
		}
	}

}
