package main

import (
	"fmt"
)

// BTreeNode represents a node in the B-tree
type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	leaf     bool
}

// BTree represents a B-tree
type BTree struct {
	root   *BTreeNode
	degree int
}

// NewBTree creates a new B-tree with the given degree
func NewBTree(degree int) *BTree {
	return &BTree{
		root:   nil,
		degree: degree,
	}
}

// Insert inserts a key into the B-tree
func (t *BTree) Insert(key int) {
	if t.root == nil {
		t.root = &BTreeNode{
			keys:     []int{key},
			children: []*BTreeNode{},
			leaf:     true,
		}
		return
	}

	if len(t.root.keys) == 2*t.degree-1 {
		newRoot := &BTreeNode{
			keys:     []int{},
			children: []*BTreeNode{t.root},
			leaf:     false,
		}
		t.splitChild(newRoot, 0)
		t.root = newRoot
	}
	t.insertNonFull(t.root, key)
}

// insertNonFull inserts a key into a non-full node
func (t *BTree) insertNonFull(node *BTreeNode, key int) {
	i := len(node.keys) - 1
	if node.leaf {
		for i >= 0 && key < node.keys[i] {
			i--
		}
		node.keys = append(node.keys[:i+1], key)
	} else {
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++
		if len(node.children[i].keys) == 2*t.degree-1 {
			t.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		t.insertNonFull(node.children[i], key)
	}
}

// splitChild splits a full child of a node
func (t *BTree) splitChild(parent *BTreeNode, index int) {
	degree := t.degree
	child := parent.children[index]
	newChild := &BTreeNode{
		leaf:     child.leaf,
		keys:     child.keys[degree:],
		children: []*BTreeNode{},
	}
	if !child.leaf {
		newChild.children = append(newChild.children, child.children[degree:]...)
		child.children = child.children[:degree]
	}
	parent.keys = append(parent.keys[:index], append([]int{child.keys[degree-1]}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*BTreeNode{newChild}, parent.children[index+1:]...)...)
	child.keys = child.keys[:degree-1]
}

// Search searches for a key in the B-tree
func (t *BTree) Search(key int) bool {
	return t.searchKey(t.root, key)
}

// searchKey searches for a key in a node recursively
func (t *BTree) searchKey(node *BTreeNode, key int) bool {
	if node == nil {
		return false
	}
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}
	if i < len(node.keys) && key == node.keys[i] {
		return true
	}
	if node.leaf {
		return false
	}
	return t.searchKey(node.children[i], key)
}

func main() {
	bt := NewBTree(3)
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys {
		bt.Insert(key)
	}

	fmt.Println("B-tree after insertion:")
	fmt.Println("Search for 6:", bt.Search(6)) // Output: true
	fmt.Println("Search for 8:", bt.Search(8)) // Output: false
}
