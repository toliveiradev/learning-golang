package main

import (
	"fmt"
	"io"
)

// Node represents a node in the binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

// createNode creates a new node with the given data
func (t *BinaryTree) createNode(data int) *Node {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t.root
}

// insert inserts a new node with the given data into the binary tree
func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func (t *BinaryTree) BFS() {
	if t.root == nil {
		return
	}
	queue := []*Node{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func main() {
	// Create a binary tree
	tree := &BinaryTree{}

	// Insert data into the tree
	tree.createNode(10)
	tree.createNode(5)
	tree.createNode(15)
	tree.createNode(3)
	tree.createNode(7)
	tree.createNode(12)
	tree.createNode(18)

	// Print the tree in in-order traversal
	fmt.Println("In-order traversal of the binary tree:")
	tree.BFS()
}
