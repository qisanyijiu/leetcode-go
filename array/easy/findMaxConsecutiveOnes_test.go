package easy

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	cases := []struct{
		arr []int
		cnt int
	}{
		struct {
			arr []int
			cnt int
		}{arr: []int{1, 1, 0, 1, 1, 1}, cnt: 3},
		struct {
			arr []int
			cnt int
		}{arr: []int{1, 0, 1}, cnt: 1},
		struct {
			arr []int
			cnt int
		}{arr: []int{1,0,1,0,1, 1,0,1}, cnt: 2},
	}
	for _, item := range cases{
		result := FindMaxConsecutiveOnes(item.arr)
		if item.cnt != result{
			t.Errorf("Arr: %v, Cnt: %d, Except: %d\n", item.arr, result, item.cnt)
		}
	}
}
