package main

import (
	"fmt"
)

// Node represents a node in the binary search tree.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// SearchResult represents the result of a search operation,
// including whether the key exists and the depth of the node.
type SearchResult struct {
	Found bool
	Depth int
}

// Insert inserts a new key into the binary search tree.
func (n *Node) Insert(key int) *Node {
	if n == nil {
		return &Node{Key: key}
	}

	if key < n.Key {
		n.Left = n.Left.Insert(key)
	} else if key > n.Key {
		n.Right = n.Right.Insert(key)
	}

	return n
}

// Search searches for a key in the binary search tree.
func (n *Node) Search(key int, depth int) SearchResult {
	if n == nil {
		return SearchResult{Found: false, Depth: -1}
	}

	if key == n.Key {
		return SearchResult{Found: true, Depth: depth}
	}

	if key < n.Key {
		return n.Left.Search(key, depth+1)
	}

	return n.Right.Search(key, depth+1)
}

func main() {
	// Create an empty binary search tree
	var root *Node

	// Insert some keys into the binary search tree
	keys := []int{10, 5, 15, 3, 7, 12, 17}
	for _, key := range keys {
		root = root.Insert(key)
	}

	// Search for keys in the binary search tree
	searchKeys := []int{3, 8, 15}
	for _, key := range searchKeys {
		result := root.Search(key, 0)
		if result.Found {
			fmt.Printf("Key %d is found in the binary search tree at depth %d.\n", key, result.Depth)
		} else {
			fmt.Printf("Key %d is not found in the binary search tree.\n", key)
		}
	}
}
