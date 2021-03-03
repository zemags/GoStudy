// Single linked list
// [a|*]->[b|*]->[x]
package main

import (
	"errors"
	"fmt"
	"log"
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
	GetAFirst() (*Node, error) // get all items except first
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
	if n.IsEmpty() {
		*n = Node{value: v, idx: size}
	} else {
		for i := 0; i < size; i++ {
			if n.next != nil {
				currentValue := n.value
				nextNodes := n.next
				fmt.Println(">", n.value, currentValue, nextNodes.value)
			}

			n = n.next
		}
		// fmt.Println(*n)
		// n.value = v
	}
	fmt.Println()
	size++
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
func (n *Node) GetAFirst() (*Node, error) {
	if !n.IsEmpty() {
		return n.next, nil
	}
	return &Node{}, errors.New("Empty list")
}

// GetLen ...
func (n *Node) GetLen() (int, error) {
	if !n.IsEmpty() {
		return size, nil
	}
	return 0, errors.New("Empty list")
}

// DisplayList ...
func DisplayList(n *Node) ([]int, error) {
	if !n.IsEmpty() {
		var s []int
		for n != nil {
			s = append(s, n.value)
			n = n.next
		}
		return s, nil
	}
	return []int{}, errors.New("empty list")
}

func main() {
	singleLL := &Node{}
	singleLL.IsEmpty()
	singleLL.AddBackword(55)
	singleLL.AddBackword(66)
	singleLL.AddBackword(77)
	singleLL.AddForward(44)
	_, err := singleLL.GetAFirst()
	if err != nil {
		log.Fatal(err)
	}
	l, _ := DisplayList(singleLL)
	fmt.Println(l)
}
