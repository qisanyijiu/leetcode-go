package lru

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LinkList struct {
	head *Node
	tail *Node
}

func NewLinkList() *LinkList {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return &LinkList{
		head: head,
		tail: tail,
	}
}

func (l *LinkList) Put(n *Node) {
	n.Next = l.head.Next
	n.Prev = l.head
	l.head.Next.Prev = n
	l.head.Next = n
}

func (l *LinkList) Remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	n.Prev = nil
}

func (l *LinkList) MoveToHead(n *Node) {
	l.Remove(n)
	l.Put(n)
}

func (l *LinkList) RemoveTail() *Node {
	tail := l.tail.Prev
	l.Remove(tail)
	return tail
}

type LRUCache struct {
	keyMap   map[int]*Node
	linkList *LinkList
	cap      int
	size     int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		keyMap:   make(map[int]*Node),
		linkList: NewLinkList(),
		cap:      cap,
		size:     0,
	}
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.keyMap[key]; ok {
		node.Value = value
		c.linkList.MoveToHead(node)
	} else {
		node := &Node{
			Key:   key,
			Value: value,
		}
		c.linkList.Put(node)
		c.keyMap[key] = node
		if c.size == c.cap {
			tail := c.linkList.RemoveTail()
			delete(c.keyMap, tail.Key)
		} else {
			c.size++
		}
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.keyMap[key]; ok {
		return node.Value
	}
	return -1
}
