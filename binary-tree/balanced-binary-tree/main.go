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

// createBalancedTree creates a balanced binary tree from sorted data
func createBalancedTree(sortedData []int) *Node {
	if len(sortedData) == 0 {
		return nil
	}

	mid := len(sortedData) / 2
	root := &Node{data: sortedData[mid]}
	root.left = createBalancedTree(sortedData[:mid])
	root.right = createBalancedTree(sortedData[mid+1:])

	return root
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
	// Sorted input data
	sortedData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a balanced binary tree
	root := createBalancedTree(sortedData)

	// Print the tree in in-order traversal
	fmt.Println("In-order traversal of the binary tree:")
	root.printInOrder()
}
