// Singly linked list
// [a|*]->[b|*]->[x]
package main

import (
	"fmt"
	"reflect"
)

// Node ...
type Node struct {
	size  int
	value int
	next  *Node
}

// SingleLinkedList ...
type SingleLinkedList interface {
	IsEmpty() bool
	AddForward(int)
	AddBackword(int)
	AddInto()
	GetFirst()
	GetAFirst() // get all items except first
	GetLen() int
}

// IsEmpty ...
func (n *Node) IsEmpty() bool {
	if reflect.DeepEqual(n, &Node{}) {
		return true
	}
	return false
}

// AddForward ...
func (n *Node) AddForward(v int) {

}

// AddBackword ...
func (n *Node) AddBackword(v int) {
	nextNode := &Node{}
	nextNode.value = v
	if n.IsEmpty() {
		*n = *nextNode
		fmt.Println(*n, n)
	} else {
		oldNode := n
		for oldNode.next != nil {
			oldNode = oldNode.next
		}
		oldNode.next = nextNode
	}
	n.size++
}

// AddInto insert by index
func (n *Node) AddInto() {

}

// GetFirst ...
func (n *Node) GetFirst() {

}

// GetAFirst ...
func (n *Node) GetAFirst() {

}

// GetLen ...
func (n *Node) GetLen() int {
	return n.size
}

func main() {
	singleLL := &Node{}
	singleLL.IsEmpty()
	singleLL.AddBackword(55)
	singleLL.AddBackword(66)
	singleLL.AddBackword(77)
	singleLL.GetLen()
}
