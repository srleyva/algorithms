package algorithms

import "fmt"

// Node in a tree
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// NewNode returns a new node in the tree
func NewNode(value int) *Node {
	return &Node{nil, nil, value}
}

// Insert inserts a new node into the list
func (n *Node) Insert(value int) {
	newNode := NewNode(value)
	if value <= n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.Insert(value)
		}
	}
}

// OrderPrint will traverse the tree recursively and print the numbers in order
func (n *Node) OrderPrint() {
	if n.Left != nil {
		n.Left.OrderPrint()
	}
	fmt.Printf("%d ", n.Value)
	if n.Right != nil {
		n.Right.OrderPrint()
	}
}

// PathWithSum tells if there is a path that adds up to passed param
func (n *Node) PathWithSum(sum int) bool {
	rval := false
	if n == nil {
		return false
	}
	sum -= n.Value

	if n.Right == nil && n.Left == nil {
		if sum == 0 {
			return true
		}
	}

	if n.Right != nil {
		if n.Right.PathWithSum(sum) {
			rval = true
		}
	}

	if n.Left != nil {
		if n.Left.PathWithSum(sum) {
			rval = true
		}
	}
	return rval
}
