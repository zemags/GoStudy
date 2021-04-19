package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(x int) *Tree {
	if t.root == nil {
		t.root = &Node{data: x, left: nil, right: nil}
	} else {
		t.root.insert(x)
	}
	return t
}

func (n *Node) insert(x int) {
	if n == nil {
		return
	} else if x <= n.data {
		if n.left != nil {
			n.left.insert(x)
		} else {
			n.left = &Node{data: x, left: nil, right: nil}
		}
	} else {
		if n.right != nil {
			n.right.insert(x)
		} else {
			n.right = &Node{data: x, left: nil, right: nil}
		}
	}
}

func NewTree() *Tree {
	return &Tree{}
}

func main() {
	t := NewTree()
	t.Insert(3)
	t.Insert(2)
	t.Insert(4)
	t.Insert(1)
	t.Insert(2)
	t.Insert(5)
	t.Insert(2)
}
