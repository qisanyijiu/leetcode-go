package medium

import "testing"

func TestSubsets(t *testing.T){
	if len(subsets([]int{1,2,3})) != 8{
		t.Fail()
	}
}
