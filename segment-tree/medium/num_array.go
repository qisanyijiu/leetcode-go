package medium

type NumArray struct {
	arr  []int
	tree []int
}

func intPow(x, y int) int {
	res := x
	for i := 1; i < y; i++ {
		res = res * x
	}
	return res
}

func buildTree(arr []int, tree *[]int, node int, start int, end int) {
	if start == end {
		(*tree)[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	nodeL := node*2 + 1
	nodeR := node*2 + 2
	buildTree(arr, tree, nodeL, start, mid)
	buildTree(arr, tree, nodeR, mid+1, end)
	(*tree)[node] = (*tree)[nodeL] + (*tree)[nodeR]
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			arr:  []int{},
			tree: []int{},
		}
	}
	treeLength := intPow(2, len(nums))
	tree := make([]int, treeLength)
	buildTree(nums, &tree, 0, 0, len(nums)-1)
	return NumArray{
		arr:  nums,
		tree: tree,
	}
}

func update(tree *[]int, node int, start, end int, val int, idx int) {
	if start == end {
		(*tree)[node] = val
		return
	}
	mid := (start + end) / 2
	nodeL := node*2 + 1
	nodeR := node*2 + 2
	if idx > mid {
		update(tree, nodeR, mid+1, end, val, idx)
	} else {
		update(tree, nodeL, start, mid, val, idx)
	}
	(*tree)[node] = (*tree)[nodeR] + (*tree)[nodeL]
}

func (this *NumArray) Update(i int, val int) {
	this.arr[i] = val
	update(&this.tree, 0, 0, len(this.arr)-1, val, i)
}

func query(tree *[]int, node int, start, end int, L, R int) int {
	if L > end || R < start {
		return 0
	} else if R >= end && L <= start {
		return (*tree)[node]
	} else if start == end {
		return (*tree)[node]
	}else {
		mid := (start + end) / 2
		nodeL := node*2 + 1
		nodeR := node*2 + 2
		sumLeft := query(tree, nodeL, start, mid, L, R)
		sumRight := query(tree, nodeR, mid+1, end, L, R)
		return sumLeft + sumRight
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return query(&this.tree, 0, 0, len(this.arr)-1, i, j)
}
