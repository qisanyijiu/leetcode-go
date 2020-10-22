package easy

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	cases := []struct{
		nums []int
		res int
	}{
		//struct {
		//	nums []int
		//	res  int
		//}{nums: []int{2, 6, 4, 8, 10, 9, 15}, res: 5},
		//struct {
		//	nums []int
		//	res  int
		//}{nums: []int{1, 3, 2,2,2}, res: 4},
		//struct {
		//	nums []int
		//	res  int
		//}{nums: []int{2,3,3,2,4}, res: 3},
		//struct {
		//	nums []int
		//	res  int
		//}{nums: []int{1,2,4,5,3}, res: 3},
		struct {
			nums []int
			res  int
		}{nums: []int{1,3,5,4,2}, res: 4},
	}
	for _, item := range cases{
		out := FindUnsortedSubarray(item.nums)
		if out != item.res{
			t.Errorf("Arr: %v, Got: %d, Except: %d\n", item.nums, out, item.res)
		}
	}
}