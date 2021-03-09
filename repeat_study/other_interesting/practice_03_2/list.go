package practice032

import "fmt"

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
func (l *LinkedList) First() interface{} {
	return l.head
}

// Last return last item
func (l *LinkedList) Last() interface{} {
	return l.tail
}

// AddForward add value at the begin of list
func (l *LinkedList) AddForward(v int) {
	l.tail = l.head.Next
	l.head.Value = v
	l.length++
	fmt.Println(l.head, l.tail)
}

// DisplayList write items from Node to list
func DisplayList(l *LinkedList) (s []interface{}) {

	// fmt.Println(l)
	// fmt.Println(l.head.Next)

	return s
}
