package easy

import "math"

type thirdTree struct {
	value int
	left  int
	right int
}

func NewThirdTree() *thirdTree{
	return &thirdTree{
		value:  math.MinInt64,
		left:   math.MinInt64,
		right:  math.MinInt64,
	}
}

func (t *thirdTree)Insert(value int){
	if value == t.right || value == t.value || value == t.left {
		return
	}
	if value > t.right{
		t.left = t.value
		t.value = t.right
		t.right = value
	} else if value > t.value{
		t.left = t.value
		t.value = value
	} else if value > t.left{
		t.left = value
	}
}

func (t *thirdTree)Full()bool{
	if t.left == math.MinInt64{
		return false
	}
	return true
}

func thirdMax(nums []int) int {
	length := len(nums)
	tree := NewThirdTree()
	for i := 0; i < length; i ++ {
		tree.Insert(nums[i])
	}
	if tree.Full(){
		return tree.left
	}else{
		return tree.right
	}
}