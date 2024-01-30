package main

import "fmt"

var count int

// Node struct represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k { // move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k { // move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k { // move right
		return n.Right.Search(k)
	} else if n.Key > k { // move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(25)
	tree.Insert(300)
	tree.Insert(400)
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(75)
	tree.Insert(12)
	tree.Insert(37)
	tree.Insert(125)
	tree.Insert(175)
	tree.Insert(110)
	fmt.Println(tree.Search(200))
	fmt.Println(count)
}
