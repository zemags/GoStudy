package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Stack Push", func(t *testing.T) {
		expected := []int{1, 2}
		actual := Stack{}
		actual.Push(1)
		actual.Push(2)
		assert.Equal(t, expected, actual.items)
	})

	t.Run("Stack Pop", func(t *testing.T) {
		expected := []int{1}
		actual := Stack{}
		actual.Push(1)
		actual.Push(2)
		actualValue := actual.Pop()
		assert.Equal(t, 2, actualValue)
		assert.Equal(t, expected, actual.items)
	})
}

func TestQueue(t *testing.T) {
	t.Run("Queue Enqueue", func(t *testing.T) {
		expected := []int{1, 2}
		actual := Queue{}
		actual.Enqueue(1)
		actual.Enqueue(2)
		assert.Equal(t, expected, actual.items)
	})

	t.Run("Queue Dequeue", func(t *testing.T) {
		expected := []int{2}
		actual := Queue{}
		actual.Enqueue(1)
		actual.Enqueue(2)
		actualValue := actual.Dequeue()
		assert.Equal(t, 1, actualValue)
		assert.Equal(t, expected, actual.items)
	})
}
