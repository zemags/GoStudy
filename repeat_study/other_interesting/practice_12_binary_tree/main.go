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
		// create new root
		t.root = &Node{data: x, left: nil, right: nil}
	} else {
		// add to existing root
		t.root.insert(x)
	}
	return t
}

func (n *Node) insert(x int) {
	if n == nil {
		return
	} else if x <= n.data {
		if n.left != nil {
			// if left child not empty
			n.left.insert(x)
		} else {
			n.left = &Node{data: x, left: nil, right: nil}
		}
	} else {
		if n.right != nil {
			// if right child not empty
			n.right.insert(x)
		} else {
			n.right = &Node{data: x, left: nil, right: nil}
		}
	}
}

func NewTree() *Tree {
	return &Tree{}
}

func printTree(t *Tree) {
	for t.root != nil {
		// fmt.Println(t.root.data)
	}
}

// remove by value
func (t *Tree) Remove(x int) *Tree {
	// find node with value
	return t
}

func (t *Tree) Contains(x int) bool {
	if t.root != nil {
		return t.root.find(x) != nil
	}
	return false
}

func (n *Node) find(x int) *Node {
	current := n
	for current != nil {
		if current != nil {
			if current.data == x {
				return current
			}
			if x < current.data {
				current = current.left
			} else {
				current = current.right
			}
		}
	}
	return current
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
