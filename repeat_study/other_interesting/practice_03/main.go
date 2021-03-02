// Single linked list
// [a|*]->[b|*]->[x]
package main

import (
	"errors"
	"fmt"
	"reflect"
)

var size int

// Node ...
type Node struct {
	idx   int
	value int
	next  *Node
}

// SingleLinkedList ...
type SingleLinkedList interface {
	IsEmpty() bool
	AddForward(int)
	AddBackword(int)
	AddInto(int, int)
	GetFirst() (int, error)
	GetAFirst() // get all items except first
	GetLen() (int, error)
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
	nextNode := Node{}
	nextNode.value = v

}

// AddBackword ...
func (n *Node) AddBackword(v int) {
	nextNode := Node{value: v, idx: size}
	if n.IsEmpty() {
		*n = nextNode
	} else {
		oldNode := n
		for oldNode.next != nil {
			oldNode = oldNode.next
		}
		oldNode.next = &nextNode
	}
	size++
}

// AddInto insert by index
func (n *Node) AddInto(v, idx int) {

}

// GetFirst ...
func (n *Node) GetFirst() (int, error) {
	if !n.IsEmpty() {
		return n.value, nil
	}
	return 0, errors.New("Empty list")
}

// GetAFirst ...
func (n *Node) GetAFirst() {
	fmt.Println(n)
}

// GetLen ...
func (n *Node) GetLen() (int, error) {
	if !n.IsEmpty() {
		return size, nil
	}
	return 0, errors.New("Empty list")
}

func main() {
	singleLL := &Node{}
	singleLL.IsEmpty()
	singleLL.AddBackword(55)
	singleLL.AddBackword(66)
	singleLL.AddBackword(77)
	singleLL.GetAFirst()
}
