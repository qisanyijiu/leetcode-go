package medium

import "testing"

func Test_CoinChange(t *testing.T) {
	type Case struct {
		Coins  []int
		Amount int
		Cnt    int
	}
	cases := []Case{
		{Coins: []int{}, Amount: 0, Cnt: 0},
		{Coins: []int{1, 2, 5}, Amount: 11, Cnt: 3},
		{Coins: []int{2}, Amount: 3, Cnt: -1},
		{Coins: []int{1}, Amount: 0, Cnt: 0},
	}
	for i, c := range cases {
		got := coinChange(c.Coins, c.Amount)
		if got != c.Cnt {
			t.Errorf("case %d: got %d, want %d", i, got, c.Cnt)
		}
	}
}
