package main

import (
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// createNode creates a new node with the given data
func createNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

// insert inserts a new node with the given data into the binary tree
func (n *Node) insert(data int, depth int, maxDepth int) {
	if depth == maxDepth {
		return
	}
	if n.left == nil {
		n.left = createNode(data)
	} else {
		n.left.insert(data, depth+1, maxDepth)
	}
	if n.right == nil {
		n.right = createNode(data)
	} else {
		n.right.insert(data, depth+1, maxDepth)
	}
}

// printInOrder prints the nodes of the tree in in-order traversal
func (n *Node) printInOrder() {
	if n == nil {
		return
	}
	n.left.printInOrder()
	fmt.Print(n.data, " ")
	n.right.printInOrder()
}

func main() {
	// Create a root node
	root := createNode(1)

	// Insert data into the tree
	maxDepth := 4 // Change the value to adjust the depth of the perfect binary tree
	for i := 2; i <= (1<<maxDepth)-1; i++ {
		root.insert(i, 1, maxDepth)
	}

	// Print the tree in in-order traversal
	fmt.Println("In-order traversal of the binary tree:")
	root.printInOrder()
}
