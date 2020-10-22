package easy

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	// 压栈
	this.data = append(this.data, x)
	// 维护最小栈
	minIndex := len(this.min) - 1
	if minIndex < 0 {
		this.min = append(this.min, x)
		return
	}
	if x <= this.min[minIndex] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	dataIndex := len(this.data) - 1
	// 如果为空栈，则直接返回
	if dataIndex < 0 {
		return
	}
	// 如果栈顶等于最小栈栈顶，则弹出最小栈栈顶
	minIndex := len(this.min) - 1
	if this.Top() == this.min[minIndex] {
		this.min = this.min[0:minIndex]
	}
	// 弹出元素
	this.data = this.data[0:dataIndex]
}

func (this *MinStack) Top() int {
	dataIndex := len(this.data) - 1
	if dataIndex < 0 {
		panic("empty stack")
	}
	return this.data[dataIndex]
}

func (this *MinStack) GetMin() int {
	minIndex := len(this.min) - 1
	if minIndex < 0{
		panic("empty stack")
	}
	return this.min[minIndex]
}
