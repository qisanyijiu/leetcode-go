package easy

import "testing"

func TestFindDisappearedNumbers(t *testing.T){
	var data []int
	data = []int{4,3,2,7,8,2,3,1}
	result := findDisappearedNumbers(data)
	if len(result) != 2 && result[1] != 5 && result[2] != 6 {
		t.Fail()
	}
}