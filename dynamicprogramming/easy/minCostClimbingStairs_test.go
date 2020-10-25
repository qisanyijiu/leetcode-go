package easy

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	cases := []struct{
		cost []int
		out int
	}{
		{cost: []int{10,15,20}, out: 15},
		{cost: []int{1,100,1,1,1,100,1,1,100,1}, out: 6},
	}
	for _, item := range cases{
		res := MinCostClimbingStairs(item.cost)
		if res != item.out {
			t.Errorf("input: %v, except: %d, got: %d", item.cost, item.out, res)
		}
	}
}
