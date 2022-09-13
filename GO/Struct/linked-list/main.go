package linkedList

import "errors"

// each node in one-way linked list know about next node.

// LinkedList struct it's just link to head-node
type LinkedList struct {
	head *Node
}

// Node struct it's data and link to next node
type Node struct {
	Data interface{}
	Next *Node
}

// Append method add new element to the end
func (l *LinkedList) Append(n *Node) error {
	if l.head == nil {
		l.head = n
		return nil
	}
	node := l.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = n
	return nil
}

// Prepend method add element to the head
func (l *LinkedList) Prepend(n *Node) {
	l.head, n.Next = n, l.head
}

// Shift method delete and return first element
func (l *LinkedList) Shift() (*Node, error) {
	if l.head == nil {
		return nil, errors.New("список пустой")
	}
	n := l.head
	l.head = l.head.Next
	return n, nil
}

// Search method return true if element exist in list and false otherwise
func (l *LinkedList) Search(d interface{}) bool {
	node := l.head
	if node == nil {
		return false
	}
	if node.Data == d {
		return true
	}
	for node.Next != nil {
		if node.Next.Data == d {
			return true
		}
		node = node.Next
	}
	return false
}

// Delete method delete first node with given data
func (l *LinkedList) Delete(d interface{}) {
	node := l.head
	if node == nil {
		return
	}
	if node.Data == d {
		l.head = node.Next
		return
	}
	for node.Next != nil {
		if node.Next.Data == d {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
	return
}
