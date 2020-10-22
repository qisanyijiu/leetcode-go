package medium

import "testing"

func TestLengthOfLongestSubString(t *testing.T){
	cases := []struct{
		S string
		Result int
	}{
		{S: "abcabcbb", Result: 3},
		{S: "bbbb", Result: 1},
	}
	for _, item := range cases{
		result := lengthOfLongestSubString(item.S)
		if result != item.Result{
			t.Errorf("String: %s,  Except: %d, Got: %d\n", item.S, item.Result, result)
		}
	}
}