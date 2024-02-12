package main

import (
	"fmt"
)

// SegmentTree represents a segment tree
type SegmentTree struct {
	tree   []int
	length int
}

// NewSegmentTree creates a new segment tree from the given array
func NewSegmentTree(arr []int) *SegmentTree {
	length := len(arr)
	tree := make([]int, 4*length)
	buildTree(arr, tree, 0, 0, length-1)
	return &SegmentTree{tree: tree, length: length}
}

// buildTree recursively builds the segment tree
func buildTree(arr, tree []int, index, left, right int) {
	if left == right {
		tree[index] = arr[left]
		return
	}
	mid := (left + right) / 2
	buildTree(arr, tree, 2*index+1, left, mid)
	buildTree(arr, tree, 2*index+2, mid+1, right)
	tree[index] = min(tree[2*index+1], tree[2*index+2])
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// QueryRange finds the minimum value in the given range [start, end]
func (st *SegmentTree) QueryRange(start, end int) int {
	if start < 0 || end >= st.length || start > end {
		return -1 // Invalid input range
	}
	return queryRangeUtil(st.tree, 0, 0, st.length-1, start, end)
}

// queryRangeUtil is a utility function for range querying
func queryRangeUtil(tree []int, index, left, right, start, end int) int {
	if start <= left && end >= right { // Total overlap
		return tree[index]
	}
	if start > right || end < left { // No overlap
		return -1 // Return some invalid value
	}
	mid := (left + right) / 2
	leftMin := queryRangeUtil(tree, 2*index+1, left, mid, start, end)
	rightMin := queryRangeUtil(tree, 2*index+2, mid+1, right, start, end)
	if leftMin == -1 {
		return rightMin
	}
	if rightMin == -1 {
		return leftMin
	}
	return min(leftMin, rightMin)
}

// Update updates the value at the given index in the array and segment tree
func (st *SegmentTree) Update(index, value int) {
	if index < 0 || index >= st.length {
		return // Invalid index
	}
	updateUtil(st.tree, 0, 0, st.length-1, index, value)
}

// updateUtil is a utility function for updating a value in the segment tree
func updateUtil(tree []int, index, left, right, idx, value int) {
	if left == right {
		tree[index] = value
		return
	}
	mid := (left + right) / 2
	if idx <= mid {
		updateUtil(tree, 2*index+1, left, mid, idx, value)
	} else {
		updateUtil(tree, 2*index+2, mid+1, right, idx, value)
	}
	tree[index] = min(tree[2*index+1], tree[2*index+2])
}

func main() {
	arr := []int{5, 2, 3, 1, 8, 7}
	st := NewSegmentTree(arr)

	fmt.Println("Minimum value in range [1, 4]:", st.QueryRange(1, 4)) // Output: 1

	st.Update(2, 4)                                                                 // Update value at index 2 to 4
	fmt.Println("Minimum value in range [1, 4] after update:", st.QueryRange(1, 4)) // Output: 1 => 1

	fmt.Println("Minimum value in range [0, 5]:", st.QueryRange(0, 5)) // Output: 1
}
