package easy

import "testing"

func TestRepeatedSubstringPattern(t *testing.T){
	cases := []struct{
		S string
		Result bool
	}{
		{S: "abc", Result: false},
		{S: "abcabc", Result: true},
	}

	for _, item := range cases{
		res := repeatedSubstringPattern(item.S)
		if res != item.Result{
			t.Errorf("String %s, Except: %v, Got: %v\n", item.S, item.Result, res)
		}
	}
}
