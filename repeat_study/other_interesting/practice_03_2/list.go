package practice032

// TwoLinkedList ...
type TwoLinkedList interface {
	Len() int
	First() *Node // get first item
	Last() *Node
	AddForward(v interface{}) *Node
	AddBackword(v interface{}) *Node
	Remove(i *Node)
	MoveToFisrt(i *Node) // move item to the first index
}

// Node ...
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// LinkedList ...
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// NewList ...
func NewList() *LinkedList {
	return &LinkedList{}
}

// Len return list length
func (l *LinkedList) Len() int {
	return l.length
}

// First return first item
func (l *LinkedList) First() *Node {
	return l.head
}

// Last return last item
func (l *LinkedList) Last() *Node {
	return l.tail
}

// AddForward add value at the begin of list
func (l *LinkedList) AddForward(v int) {
	node := &Node{}
	node.Value = v
	node.Next = l.head
	node.Prev = nil

	if l.head != nil {
		l.head.Prev = node
	}
	if l.tail == nil {
		l.tail = node
	}
	l.head = node
	l.length++
}

// AddBackword add value at the begin of list
func (l *LinkedList) AddBackword(v int) {
	node := &Node{}
	node.Value = v
	node.Next = nil
	node.Prev = l.tail

	if l.tail != nil {
		// if item was in list tail
		l.tail.Next = node
	}

	l.tail = node
	if l.head == nil {
		l.head = node
	}

	l.length++
}

// DisplayList write items from Node to list
func DisplayList(l *LinkedList) (s []interface{}) {
	for i := 0; i < l.length; i++ {
		s = append(s, l.head.Value)
		l.head = l.head.Next
	}
	return s
}
