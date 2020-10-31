package hard

import (
	"leetcode/utils"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	type CaseItem struct {
		S      string
		Words  []string
		Except []int
	}
	cases := []CaseItem{
		{
			S:      "barfoothefoobarman",
			Words:  []string{"foo","bar"},
			Except: []int{0,9},
		},
		{
			S:      "wordgoodgoodgoodbestword",
			Words:  []string{"word","good","best","word"},
			Except: []int{},
		},
	}
	for _, item := range cases{
		res := FindSubstring(item.S, item.Words)
		if !utils.EqualIntSlice(item.Except, res) {
			t.Errorf("S: %s, Words: %v, except: %v, got: %v", item.S, item.Words, item.Except, res)
		}
	}
}