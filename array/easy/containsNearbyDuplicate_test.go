package easy

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	cases := []struct{
		nums []int
		k int
		res bool
	}{
		struct {
			nums []int
			k int
			res  bool
		}{nums: []int{1,2,3,1}, k: 3,  res: true},
		struct {
			nums []int
			k int
			res  bool
		}{nums: []int{1,0,1, 1}, k: 1, res: true},
		struct {
			nums []int
			k int
			res  bool
		}{nums: []int{1,2,3,1,2,3},k: 2, res: false},
		struct {
			nums []int
			k int
			res  bool
		}{nums: []int{99, 99},k: 2, res: true},
		struct {
			nums []int
			k int
			res  bool
		}{nums: []int{1, 2},k: 2, res: false},
	}
	for _, item := range cases{
		res := ContainsNearbyDuplicate(item.nums, item.k)
		if res != item.res{
			t.Errorf("Arr: %v, k: %d, Got: %t, Except: %t\n", item.nums, item.k, res, item.res)
		}
	}
}