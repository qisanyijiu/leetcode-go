package easy

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct{
		nums []int
		miss int
	}{
		struct {
			nums []int
			miss int
		}{nums: []int{3,0,1}, miss: 2},
		struct {
			nums []int
			miss int
		}{nums: []int{9,6,4,2,3,5,7,0,1}, miss: 8},
	}
	for _, item := range cases{
		res := MissingNumber(item.nums)
		if res != item.miss{
			t.Errorf("Arr: %v, Got: %d, Except: %d\n", item.nums, res, item.miss)
		}
	}
}
