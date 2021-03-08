package practice032

// two linked list

// TwoLinkedList ...
type TwoLinkedList interface {
	Len() int
	First() *Node
	Last() *Node
	AddForward(v interface{}) *Node
	AddBackword(v interface{}) *Node
	Remove(i *Node)
	MoveToFront(i *Node)
}

// Node ...
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type linkedList struct {
	TwoLinkedList
}

// NewList ...
func NewList() TwoLinkedList {
	return new(linkedList)
}
