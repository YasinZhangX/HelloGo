package tree

import "fmt"

// Node in binary tree
type Node struct {
	Value       int
	Left, Right *Node
}

// CreateNode create new tree node
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// Print print node value
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue set node value
func (node *Node) SetValue(value int) {
	node.Value = value
}
