package easy

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct{
		nums []int
		res bool
	}{
		struct {
			nums []int
			res  bool
		}{nums: []int{1,2,3,1}, res: true},
		struct {
			nums []int
			res  bool
		}{nums: []int{1,2,3,4}, res: false},
		struct {
			nums []int
			res  bool
		}{nums: []int{1,2,3,1,1}, res: true},
	}
	for _, item := range cases{
		res := ContainsDuplicate(item.nums)
		if res != item.res{
			t.Errorf("Arr: %v, Got: %t, Except: %t\n", item.nums, res, item.res)
		}
	}
}
