package easy

import "testing"

func TestThirdMax(t *testing.T) {
	var data []int
	data = []int{3, 2, 1}
	if thirdMax(data) != 1{
		t.Fail()
	}
	data = []int{1, 2}
	if thirdMax(data) != 1{
		t.Fail()
	}
	data = []int{3, 2, 2,1}
	if thirdMax(data) != 1{
		t.Fail()
	}
	data = []int{2,2,3,1}
	if thirdMax(data) != 1{
		t.Fail()
	}
	data = []int{1,2,2,5,3,5}
	if thirdMax(data) != 2{
		t.Fail()
	}
}