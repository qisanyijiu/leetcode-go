package medium

import "testing"

func TestEquationsPossible(t *testing.T) {
	cases := []struct{
		Input []string
		Res bool
	}{
		{Input: []string{"a==b","b!=a"}, Res: false},
		{Input: []string{"b==a","a==b"}, Res:true},
		{Input: []string{"a==b","b==c","a==c"}, Res:true},
		{Input: []string{"a==b","b!=c","c==a"}, Res:false},
		{Input: []string{"c==c","b==d","x!=z"}, Res:true},
	}
	for _, item := range cases {
		out := EquationsPossible(item.Input)
		if out != item.Res {
			t.Errorf("input: %s, except: %t, got: %t", item.Input,item.Res, out)
		}
	}
}