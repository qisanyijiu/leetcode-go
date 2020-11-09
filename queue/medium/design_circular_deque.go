package medium

/**
641. 设计循环双端队列

https://leetcode-cn.com/problems/design-circular-deque/

设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。
 */
type MyCircularDeque struct {
	head  *Node
	tail  *Node
	len   int
	count int
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	head.Next = &tail
	tail.Pre = &head
	deque := MyCircularDeque{
		head:  &head,
		tail:  &tail,
		len:   k,
		count: 0,
	}

	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.head.Next
	tempNode := Node{
		Next: temp,
		Pre:  this.head,
		Val:  value,
	}
	this.head.Next = &tempNode
	temp.Pre = &tempNode
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.tail.Pre
	tempNode := Node{
		Next: this.tail,
		Pre:  temp,
		Val:  value,
	}
	this.tail.Pre = &tempNode
	temp.Next = &tempNode
	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.head.Next
	this.head.Next = deleteTemp.Next
	deleteTemp.Next.Pre = this.head
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.tail.Pre
	deleteTemp.Pre.Next = this.tail
	this.tail.Pre = deleteTemp.Pre
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head.Next == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.count
}