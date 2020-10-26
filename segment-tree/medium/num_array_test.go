package medium

import "testing"

func TestNumArray(t *testing.T) {
	numArray := Constructor([]int{1,3,5})
	if numArray.SumRange(0,2) != 9 {
		t.Errorf("except: 9, got: %d", numArray.SumRange(0,2))
	}
	numArray.Update(1,2)
	if numArray.SumRange(0,2) != 8 {
		t.Errorf("except: 8, got: %d", numArray.SumRange(0,2))
	}
}

func TestNumArray1(t *testing.T) {
	numArray := Constructor([]int{9,-8})
	numArray.Update(0,3)
	if numArray.SumRange(1,1) != -8 {
		t.Errorf("except: -8, got: %d", numArray.SumRange(0,2))
	}
	if numArray.SumRange(0,1) != -5 {
		t.Errorf("except: -5, got: %d", numArray.SumRange(0,2))
	}
	numArray.Update(1,-3)
	if numArray.SumRange(0,1) != 0 {
		t.Errorf("except: 0, got: %d", numArray.SumRange(0,2))
	}
}