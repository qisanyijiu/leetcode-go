package easy

import "testing"


func TestTwoSum(t *testing.T) {
	type Item struct {
		nums []int
		target int
		out []int
	}
	cases := []Item{
		{
			nums: []int{2, 7, 11, 15},
			target: 9,
			out:  []int{1,2},
		},
	}
	for _, item := range cases {
		ans := TwoSum(item.nums, item.target)
		if ans[0] != item.out[0] || ans[1] != item.out[1] {
			t.Errorf("nums: %v, target: %d, except: %v, got: %v", item.nums, item.target, item.out, ans)
		}
	}
}