package medium

import "testing"

func arrayEqual(a , b []int) bool {
	if len(a) != len(b){
		return false
	}
	for i, item := range a{
		if item != b[i]{
			return false
		}
	}
	return true
}

func TestSortColors(t *testing.T){
	x := []int{0,1,2,0,1,2}
	sortColors(x)
	if !arrayEqual(x, []int{0,0,1,1,2,2}){
		t.Fail()
	}
}
