// Single linked list
// [a|*]->[b|*]->[x|nil]
package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var size int

// Node implement list item
type Node struct {
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

// IsEmpty check if list is empty
func (n *Node) IsEmpty() bool {
	if reflect.DeepEqual(n, &Node{}) {
		return true
	}
	return false
}

// AddForward add item to the first place and shift other right
func (n *Node) AddForward(v int) {
	if n.IsEmpty() {
		*n = Node{value: v}
	} else {
		nextNodes := *n
		n.value = v
		n.next = &nextNodes
	}
	size++
}

// AddBackword add item to the end of list
func (n *Node) AddBackword(v int) {
	nextNode := Node{value: v}
	if n.IsEmpty() {
		*n = nextNode
	} else {
		currentNode := n
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &nextNode
	}
	size++
}

// AddInto insert by index
func (n *Node) AddInto(v, idx int) error {
	if idx == 0 && n.IsEmpty() {
		n.AddBackword(v)
	} else if idx != 0 && n.IsEmpty() || idx > size {
		return errors.New("index out or range")
	} else {
		for i := 0; i < size; i++ {
			if i == idx {
				if n.next == nil {
					nextNodes := *n
					n.next = &nextNodes
				} else {
					// TODO: refactor cause TestAddInto
					nextNodes := *n
					n.next = &nextNodes
				}
				n.value = v
			}
			n = n.next
		}
	}

	size++
	return nil
}

// GetFirst item from list
func (n *Node) GetFirst() (int, error) {
	if !n.IsEmpty() {
		return n.value, nil
	}
	return 0, errors.New("Empty list")
}

// GetAFirst other items except first from list
func (n *Node) GetAFirst() (*Node, error) {
	if !n.IsEmpty() {
		return n.next, nil
	}
	return &Node{}, errors.New("Empty list")
}

// GetLen return len of list
func (n *Node) GetLen() (int, error) {
	if !n.IsEmpty() {
		return size, nil
	}
	return 0, errors.New("Empty list")
}

// DisplayList write items from Node to list
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
	singleLL.AddBackword(88)
	_, err := singleLL.GetAFirst()
	if err != nil {
		log.Fatal(err)
	}
	singleLL.AddInto(44, 3)
	l, _ := DisplayList(singleLL)
	fmt.Println(l)
}
