package medium

import "testing"

func TestMaxScore(t *testing.T) {
	type Item struct {
		CardPoints []int
		K          int
		Res        int
	}
	cases := []Item{
		{
			CardPoints: []int{1, 2, 3, 4, 5, 6, 1},
			K:          3,
			Res:        12,
		},
		{
			CardPoints: []int{2, 2, 2},
			K:          2,
			Res:        4,
		},
		{
			CardPoints: []int{9, 7, 7, 9, 7, 7, 9},
			K:          7,
			Res:        55,
		},
		{
			CardPoints: []int{1, 1000, 1},
			K: 1,
			Res: 1,
		},
		{
			CardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
			K: 3,
			Res: 202,
		},
	}
	for _, item := range cases {
		out := MaxScore(item.CardPoints, item.K)
		if out != item.Res {
			t.Errorf("input: %v, k: %d,  except: %d, got: %d", item.CardPoints, item.K, item.Res, out)
		}
	}
}
