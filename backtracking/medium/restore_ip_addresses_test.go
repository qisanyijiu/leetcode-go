package medium

import (
	"leetcode/utils"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	cases := []struct{
		s string
		ans []string
	}{
		{s: "25525511135", ans: []string{"255.255.11.135","255.255.111.35"}},
	}
	for _, item := range cases {
		output := RestoreIpAddresses(item.s)
		if len(output) != len(item.ans) {
			t.Errorf("error ip cnt: %d, except: %d", len(output), len(item.ans))
		}else{
			for _, ip := range output {
				if !utils.InStrings(ip, item.ans) {
					t.Errorf("error ip: %s", ip)
				}
			}
		}
	}
	
}