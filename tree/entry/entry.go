package main

import (
	"fmt"

	"yasinzhang.top/hellogo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

type embedTreeNode struct {
	*tree.Node
}

func (node *embedTreeNode) traverse() {
	fmt.Println("Traverse")
	fmt.Println(node.Left)
	fmt.Println(node.Right)
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	node := myTreeNode{myNode.node.Left}
	node.postOrder()
	node = myTreeNode{myNode.node.Right}
	node.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = *tree.CreateNode(3)
	root.Left = tree.CreateNode(1)
	root.Right = tree.CreateNode(5)

	root.Print()
	root.SetValue(100)
	root.Print()

	fmt.Println("postOrder")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	fmt.Println("Traverse")
	root.SetValue(3)
	root.Traverse()

	fmt.Println("count node")
	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)

	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value: ", maxNodeValue)
}
