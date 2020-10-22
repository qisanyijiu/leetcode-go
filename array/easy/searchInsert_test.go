package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct{
		nums []int
		target, result int
	}{
		struct {
			nums           []int
			target, result int
		}{nums: []int{1,3,5,6}, target: 5, result: 2},
		struct {
			nums           []int
			target, result int
		}{nums: []int{1,3,5,6}, target: 2, result: 1},
		struct {
			nums           []int
			target, result int
		}{nums: []int{1,3,5,6}, target: 7, result: 4},
		struct {
			nums           []int
			target, result int
		}{nums: []int{1,3,5,6}, target: 0, result: 0},
		struct {
			nums           []int
			target, result int
		}{nums: []int{1}, target: 0, result: 0},
	}
	for _, item := range cases{
		result := SearchInsert(item.nums, item.target)
		if result != item.result {
			t.Errorf("Nums: %v, Target: %d, Except: %d, Got: %v\n", item.nums, item.target, item.result, result)
		}
	}
}