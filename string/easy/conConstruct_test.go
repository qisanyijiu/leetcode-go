package easy

import "testing"

func TestConConstruct(t *testing.T){
	cases := []struct{
		ransomNote string
		magazine string
		isCan bool
	}{
		{ransomNote: "a", magazine:"b", isCan:false},
		{ransomNote: "a", magazine:"ab", isCan:true},
		{ransomNote: "aab", magazine:"aaab", isCan:true},
		{ransomNote: "aabb", magazine:"aaab", isCan:false},
	}
	for _, item :=range cases{
		res := conConstruct(item.ransomNote, item.magazine)
		if res != item.isCan{
			t.Errorf("RansomNote %s, Magazine: %s, Except: %v, Got: %v\n", item.ransomNote, item.magazine, item.isCan, res)
		}
	}
}