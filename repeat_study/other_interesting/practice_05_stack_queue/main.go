package main

// Stack hold a slice
type Stack struct {
	items []int
}

// Push item to stuck
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop take item from stack and return it
func (s *Stack) Pop() int {
	l := len(s.items)
	toRemove := s.items[l-1]
	s.items = s.items[:l-1]
	return toRemove
}

// Queue
type Queue struct {
	items []int
}

// Enqueue add item to queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue remove item from queue and return it
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {

}
