package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 17}
	node3 := &node{data: 12}
	node4 := &node{data: 7}
	node5 := &node{data: 87}
	node6 := &node{data: 8}
	node7 := &node{data: 10}
	node8 := &node{data: 36}
	node9 := &node{data: 51}
	node10 := &node{data: 3}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.prepend(node10)
	myList.printListData()
	myList.deleteWithValue(99)
	myList.deleteWithValue(87)
	myList.deleteWithValue(3)
	myList.deleteWithValue(48)
	myList.printListData()
}
