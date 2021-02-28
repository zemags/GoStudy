// Singly linked list
// [a|*]->[b|*]->[x]
package main

import (
	"reflect"
)

// SingleLinkedStruct ...
type SingleLinkedStruct struct {
	value int
	next  *SingleLinkedStruct
}

// SingleLinkedList ...
type SingleLinkedList interface {
	IsEmpty() bool
	AddForward()
	AddBackword()
	AddInto()
	GetFirst()
	GetAFirst() // get all items except first
}

// IsEmpty ...
func (s *SingleLinkedStruct) IsEmpty() bool {
	if reflect.DeepEqual(s, &SingleLinkedStruct{}) {
		return true
	}
	return false
}

// AddForward ...
func (s *SingleLinkedStruct) AddForward() {

}

// AddBackword ...
func (s *SingleLinkedStruct) AddBackword() {

}

// AddInto ...
func (s *SingleLinkedStruct) AddInto() {

}

// GetFirst ...
func (s *SingleLinkedStruct) GetFirst() {

}

// GetAFirst ...
func (s *SingleLinkedStruct) GetAFirst() {

}

func main() {
	singleLL := &SingleLinkedStruct{}
	singleLL.IsEmpty()
}
