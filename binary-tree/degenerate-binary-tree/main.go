package main

import (
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	data  int
	right *Node // For the degenerate tree, we only use the right child
}

// createNode creates a new node with the given data
func createNode(data int) *Node {
	return &Node{data: data, right: nil}
}

// insert inserts a new node with the given data into the binary tree
func (n *Node) insert(data int) {
	if n.right == nil {
		n.right = createNode(data)
	} else {
		// Insert to the right recursively
		n.right.insert(data)
	}
}

// printInOrder prints the nodes of the tree in in-order traversal
func (n *Node) printInOrder() {
	if n == nil {
		return
	}
	fmt.Print(n.data, " ")
	n.right.printInOrder()
}

func main() {
	// Create a root node
	root := createNode(1)

	// Insert data into the tree to make it degenerate
	for i := 2; i <= 10; i++ {
		root.insert(i)
	}

	// Print the tree in in-order traversal
	fmt.Println("In-order traversal of the degenerate binary tree:")
	root.printInOrder()
}
