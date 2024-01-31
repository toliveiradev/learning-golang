package main

import (
	"container/heap"
	"fmt"
)

// Node represents a node in the Huffman tree
type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

// HuffmanHeap is a min-heap to build the Huffman tree
type HuffmanHeap []*Node

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0 : n-1]
	return node
}

// BuildHuffmanTree builds the Huffman tree from a min-heap
func BuildHuffmanTree(freqMap map[rune]int) *Node {
	var pq HuffmanHeap

	// Initialize the heap with leaf nodes
	for char, f := range freqMap {
		pq = append(pq, &Node{Char: char, Freq: f})
	}

	heap.Init(&pq)

	// Build the Huffman tree
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		// Create a new node with the sum of the frequencies
		// and set the left and right nodes
		internalNode := &Node{Char: 0, Freq: left.Freq + right.Freq, Left: left, Right: right}
		heap.Push(&pq, internalNode)
	}

	return heap.Pop(&pq).(*Node)
}

// CountFrequency counts the frequency of each character in a text
func CountFrequency(text string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}
	return freqMap
}

func main() {
	// Example text
	text := "this is an example for huffman encoding"

	// Count the frequency of each character
	freqMap := CountFrequency(text)

	// Build the Huffman tree
	root := BuildHuffmanTree(freqMap)

	// Print the Huffman tree
	fmt.Println("Huffman tree:")
	printHuffmanTree(root, "")

	// Print the frequency of each character
	fmt.Println("\nFrequency of each character:")
	for char, freq := range freqMap {
		fmt.Printf("%s: %d\n", string(char), freq)
	}
}

// printHuffmanTree prints the visual representation of the Huffman tree
func printHuffmanTree(node *Node, prefix string) {
	if node != nil {
		isLeaf := node.Left == nil && node.Right == nil
		fmt.Printf("%s%s (%d)\n", prefix, string(node.Char), node.Freq)
		if isLeaf {
			return
		}
		printHuffmanTree(node.Left, prefix+"| ")
		printHuffmanTree(node.Right, prefix+"| ")
	}
}
