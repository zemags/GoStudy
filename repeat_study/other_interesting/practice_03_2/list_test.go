package practice032

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()
		assert.Equal(t, 0, l.Len())
	})
}

func TestAddForward(t *testing.T) {
	expected := []interface{}([]interface{}{4, 3, 2, 1})
	l := NewList()
	for _, i := range []int{1, 2, 3, 4} {
		l.AddForward(i)
	}
	actual := DisplayList(l)
	assert.Equal(t, expected, actual)
}

func TestAddBackword(t *testing.T) {
	expected := []interface{}([]interface{}{1, 2, 3, 4})
	l := NewList()
	for _, i := range []int{1, 2, 3, 4} {
		l.AddBackword(i)
	}
	actual := DisplayList(l)
	assert.Equal(t, expected, actual)
}

func TestDisplayList(t *testing.T) {
	expected := []interface{}([]interface{}{1, 2, 3})
	l := NewList()

	l.length = 3
	l.head = &Node{
		Value: 1,
		Prev:  nil,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}

	actual := DisplayList(l)
	assert.Equal(t, expected, actual)
}

func TestFirst(t *testing.T) {
	l := NewList()
	l.AddForward(1)
	l.AddForward(2)
	l.AddForward(3)
	node := l.First()
	actual := node.Value.(int)
	assert.Equal(t, 3, actual)
}

func TestLast(t *testing.T) {
	l := NewList()
	l.AddForward(1)
	l.AddForward(2)
	l.AddForward(3)
	node := l.Last()
	actual := node.Value.(int)
	assert.Equal(t, 1, actual)
}
