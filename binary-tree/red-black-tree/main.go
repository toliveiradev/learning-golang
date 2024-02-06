package main

import (
	"fmt"
)

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type Node struct {
	key    int
	color  Color
	parent *Node
	left   *Node
	right  *Node
}

type RedBlackTree struct {
	root *Node
}

func NewNode(key int) *Node {
	return &Node{
		key:    key,
		color:  RED,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		root: nil,
	}
}

func (t *RedBlackTree) insertFixup(z *Node) {
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			uncle := z.parent.parent.right
			if uncle != nil && uncle.color == RED {
				z.parent.color = BLACK
				uncle.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			uncle := z.parent.parent.left
			if uncle != nil && uncle.color == RED {
				z.parent.color = BLACK
				uncle.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RedBlackTree) Insert(key int) {
	z := NewNode(key)
	var y *Node
	x := t.root

	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	t.insertFixup(z)
}

func (t *RedBlackTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree) rightRotate(y *Node) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

func inorderTraversal(node *Node) {
	if node != nil {
		inorderTraversal(node.left)
		fmt.Printf("%d ", node.key)
		inorderTraversal(node.right)
	}
}

func main() {
	tree := NewRedBlackTree()
	keys := []int{7, 3, 18, 10, 22, 8, 11, 26}
	for _, key := range keys {
		tree.Insert(key)
	}

	fmt.Println("Inorder Traversal:")
	inorderTraversal(tree.root)
	fmt.Println()
}
