package easy

import "testing"

func TestMerge(t *testing.T) {
	cases := []struct{
		nums1, nums2, output []int
		m, n  int
	}{
		struct {
			nums1, nums2, output []int
			m, n         int
		}{nums1: []int{1,2,3,0,0,0}, nums2: []int{2,5,6}, m: 3, n: 3},
	}
	for _, item := range cases{
		nums1 := item.nums1
		Merge(item.nums1, item.m, item.nums2, item.n)
		if !ArrEqual(item.nums1, item.output){
			t.Errorf("Nums1; %v, Nums2: %v, M: %d, N: %d, Except: %v, Got: %v\n", nums1, item.nums2, item.m, item.n, item.output, item.nums1)
		}
	}
}

func ArrEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2){
		return false
	}
	for index , item := range nums1{
		if nums2[index] != item{
			return false
		}
	}
	return true
}
