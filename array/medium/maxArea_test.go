package medium

import "testing"

func TestMaxArea(t *testing.T){
	var data []int
	data = []int{1,8,6,2,5,4,8,3,7}
	if maxArea(data) != 49 {
		t.Fail()
	}
}