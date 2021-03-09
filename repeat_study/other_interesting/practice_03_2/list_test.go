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
	l := NewList()
	l.AddForward(1)
	l.AddForward(2)
	l.AddForward(3)
}

func TestDisplayListValid(t *testing.T) {
	expected := []int{1, 2, 3}
	l := NewList()
	for _, i := range []int{1, 2, 3} {
		l.AddForward(i)
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
