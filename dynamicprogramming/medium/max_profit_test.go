package medium

import "testing"

func TestMaxProfit(t *testing.T) {
	type Item struct {
		Prices []int
		Profit int
	}
	cases := []Item{
		{
			Prices: []int{7,1,5,3,6,4},
			Profit: 7,
		},
		{
			Prices: []int{1,2,3,4,5},
			Profit: 4,
		},
		{
			Prices: []int{7,6,4,3,1},
			Profit: 4,
		},
	}
	for _, item := range cases {
		out := MaxProfit(item.Prices)
		if out != item.Profit {
			t.Errorf("input: %v, except: %d, got: %d", item.Prices, item.Profit, out)
		}
	}
}

