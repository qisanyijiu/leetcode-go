package easy

import "testing"

func TestCheckPossibility(t *testing.T){
	var data []int
	data = []int{4,1,2}
	if !checkPossibility(data){
		t.Fail()
	}
	data = []int{4,2,1}
	if checkPossibility(data){
		t.Fail()
	}
	data = []int{3,4,2,3}
	if checkPossibility(data){
		t.Fail()
	}
}