package easy

import "testing"

func TestMaxProfit(t *testing.T) {
	type Item struct {
		Prices []int
		Profit int
	}
	cases := []Item{
		{
			Prices: []int{7,1,5,3,6,4},
			Profit: 5,
		},
		{
			Prices: []int{7,6,4,3,1},
			Profit: 0,
		},
	}
	for _, item := range cases {
		out := MaxProfit(item.Prices)
		if out != item.Profit {
			t.Errorf("input: %v, except: %d, got: %d", item.Prices, item.Profit, out)
		}
	}
}
