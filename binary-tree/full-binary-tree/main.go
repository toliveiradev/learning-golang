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
func (n *Node) insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = createNode(data)
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = createNode(data)
		} else {
			n.right.insert(data)
		}
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
	root := createNode(10)

	// Insert data into the tree
	root.insert(5)
	root.insert(15)
	root.insert(3)
	root.insert(7)
	root.insert(12)
	root.insert(18)

	// Print the tree in in-order traversal
	fmt.Println("In-order traversal of the binary tree:")
	root.printInOrder()
}
