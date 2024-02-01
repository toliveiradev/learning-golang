package main

import (
	"fmt"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Height: 1}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	} else {
		return root // Duplicates are not allowed
	}

	updateHeight(root)

	balance := balanceFactor(root)

	// Left Heavy
	if balance > 1 {
		// Right rotation
		if key < root.Left.Key {
			return rightRotate(root)
		}
		// Left-Right rotation
		if key > root.Left.Key {
			root.Left = leftRotate(root.Left)
			return rightRotate(root)
		}
	}

	// Right Heavy
	if balance < -1 {
		// Left rotation
		if key > root.Right.Key {
			return leftRotate(root)
		}
		// Right-Left rotation
		if key < root.Right.Key {
			root.Right = rightRotate(root.Right)
			return leftRotate(root)
		}
	}

	return root
}

func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d ", root.Key)
		inorderTraversal(root.Right)
	}
}

func main() {
	var root *Node

	keys := []int{10, 20, 30, 40, 50, 25}

	for _, key := range keys {
		root = insert(root, key)
	}

	fmt.Println("Inorder Traversal of AVL tree:")
	inorderTraversal(root)
}
