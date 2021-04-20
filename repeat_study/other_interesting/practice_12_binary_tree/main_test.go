package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	expected := &Tree{
		root: &Node{
			data: 2,
			left: &Node{
				data:  1,
				left:  nil,
				right: nil,
			},
			right: &Node{
				data:  3,
				left:  nil,
				right: nil,
			},
		},
	}

	tree := NewTree()
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	assert.Equal(t, expected, tree)
}

func TestContains(t *testing.T) {
	tree := NewTree()
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	expectedTrue := tree.Contains(3)
	assert.Equal(t, expectedTrue, true)

	expectedFalse := tree.Contains(4)
	assert.Equal(t, expectedFalse, false)

}
