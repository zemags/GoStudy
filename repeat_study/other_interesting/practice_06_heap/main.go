package main

import "fmt"

// MaxHeap struct
type MaxHeap struct {
	array []int
}

// Insert
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array)
	if l == 0 {
		return -1 // empty heap
	}
	h.array[0] = h.array[l-1]
	h.array = h.array[:l-1]
	return extracted
}

// maxHeapifyUp will heapify from buttom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown
func (h *MaxHeap) maxHeapifyDown(index int) {
	// loop while index has as least one child
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		// if left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		//compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left
func left(i int) int {
	return i*2 + 1
}

// right
func right(i int) int {
	return i*2 + 1
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{1, 2, 3, 5, 2, 7, 56, 64, 23, 76, 34, 79, 22}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}
