package medium

import "testing"

func TestLeastInterval(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	if leastInterval(tasks, 2) !=8{
		t.Fail()
	}
}
