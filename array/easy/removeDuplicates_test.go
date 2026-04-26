package easy

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var data []int
	data = []int{1, 1, 2}
	if removeDuplicates(data) != 2 {
		t.Errorf("remove duplicates failed, expected 2, got %d", removeDuplicates(data))
	}
	data = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	if removeDuplicates(data) != 5 {
		t.Errorf("remove duplicates failed, expected 5, got %d", removeDuplicates(data))
	}
}
