package main

import (
	"fmt"
)

// BPlusTreeNode represents a node in the B+ tree
type BPlusTreeNode struct {
	keys     []int
	children []*BPlusTreeNode
	parent   *BPlusTreeNode
	leaf     bool
}

// BPlusTree represents a B+ tree
type BPlusTree struct {
	root   *BPlusTreeNode
	degree int
}

// NewBPlusTree creates a new B+ tree with the given degree
func NewBPlusTree(degree int) *BPlusTree {
	return &BPlusTree{
		root:   nil,
		degree: degree,
	}
}

// Insert inserts a key into the B+ tree
func (t *BPlusTree) Insert(key int) {
	if t.root == nil {
		t.root = &BPlusTreeNode{
			keys:     []int{key},
			children: []*BPlusTreeNode{},
			parent:   nil,
			leaf:     true,
		}
		return
	}

	if len(t.root.keys) == 2*t.degree-1 {
		newRoot := &BPlusTreeNode{
			keys:     []int{},
			children: []*BPlusTreeNode{t.root},
			parent:   nil,
			leaf:     false,
		}
		t.splitChild(newRoot, 0)
		t.root = newRoot
	}
	t.insertNonFull(t.root, key)
}

// insertNonFull inserts a key into a non-full node
func (t *BPlusTree) insertNonFull(node *BPlusTreeNode, key int) {
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
func (t *BPlusTree) splitChild(parent *BPlusTreeNode, index int) {
	degree := t.degree
	child := parent.children[index]
	newChild := &BPlusTreeNode{
		leaf:     child.leaf,
		keys:     child.keys[degree:],
		children: []*BPlusTreeNode{},
		parent:   parent,
	}
	if !child.leaf {
		newChild.children = append(newChild.children, child.children[degree:]...)
		for _, c := range newChild.children {
			c.parent = newChild
		}
		child.children = child.children[:degree]
	}
	parent.keys = append(parent.keys[:index], append([]int{child.keys[degree-1]}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*BPlusTreeNode{newChild}, parent.children[index+1:]...)...)
	child.keys = child.keys[:degree-1]
}

// Search searches for a key in the B+ tree
func (t *BPlusTree) Search(key int) bool {
	return t.searchKey(t.root, key)
}

// searchKey searches for a key in a node recursively
func (t *BPlusTree) searchKey(node *BPlusTreeNode, key int) bool {
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
	bt := NewBPlusTree(3)
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys {
		bt.Insert(key)
	}

	fmt.Println("B+ tree after insertion:")
	fmt.Println("Search for 6:", bt.Search(6)) // Output: true
	fmt.Println("Search for 8:", bt.Search(8)) // Output: false
}
