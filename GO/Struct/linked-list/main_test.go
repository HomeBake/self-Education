package linkedList

import (
	"testing"
)

func TestLinkedList_Prepend(t *testing.T) {
	var l LinkedList
	n := &Node{
		Data: "new node",
	}
	t.Run("Linked list is empty", func(t *testing.T) {
		l.Prepend(n)
		if l.head.Data != n.Data {
			t.Errorf("Expected node data %s and returned %s", n.Data, l.head.Data)
		}
		if l.head.Next != n.Next {
			t.Errorf("Expected next node %s and returned %s", n.Next, l.head.Next)
		}
	})
	t.Run("linked list not empty", func(t *testing.T) {
		l.Prepend(n)
		if l.head.Data != n.Data {
			t.Errorf("Expected node data %s and returned %s", n.Data, l.head.Data)
		}
		if l.head.Next != n.Next {
			t.Errorf("Expected next node %s and returned %s", n.Next, l.head.Next)
		}
	})
}

func TestLinkedList_Append(t *testing.T) {
	var l LinkedList
	t.Run("add to empty list", func(t *testing.T) {
		nodeData := "head element"
		expectedNode := &Node{
			Data: nodeData,
			Next: nil,
		}
		l.Append(expectedNode)
		if l.head.Data != expectedNode.Data {
			t.Errorf("Expected node data %s and returned %s", expectedNode.Data, l.head.Data)
		}
		if l.head.Next != expectedNode.Next {
			t.Errorf("Expected next node %s and returned %s", expectedNode.Next, l.head.Next)
		}
	})
	t.Run("add to only head list", func(t *testing.T) {
		nodeData := "head element"
		expectedNode := &Node{
			Data: nodeData,
		}
		l.Append(expectedNode)
		if l.head.Next.Data != expectedNode.Data {
			t.Errorf("Expected node data %s and returned %s", expectedNode.Data, l.head.Data)
		}
		if l.head.Next.Next != expectedNode.Next {
			t.Errorf("Expected next node %s and returned %s", expectedNode.Next, l.head.Next)
		}
	})
	t.Run("add to filled list", func(t *testing.T) {
		nodeData := "third element"
		expectedNode := &Node{
			Data: nodeData,
		}
		l.Append(expectedNode)
		if l.head.Next.Next.Data != expectedNode.Data {
			t.Errorf("Expected node data %s and returned %s", expectedNode.Data, l.head.Data)
		}
		if l.head.Next.Next.Next != expectedNode.Next {
			t.Errorf("Expected next node %s and returned %s", expectedNode.Next, l.head.Next)
		}
	})
}

func TestLinkedList_Shift(t *testing.T) {
	var l LinkedList
	t.Run("Linked list is empty", func(t *testing.T) {
		_, err := l.Shift()
		if err == nil {
			t.Error("Expected error but return nil")
		}
	})
	t.Run("Linked list has only head", func(t *testing.T) {
		nodeData := "head element"
		n := &Node{
			Data: nodeData,
		}
		l.Prepend(n)
		node, err := l.Shift()
		if err != nil {
			t.Error("Error returned, but not expected")
		}
		if node != n {
			t.Errorf("Returned %s but %s expected", node.Data, n.Data)
		}
		if l.head != nil {
			t.Error("Expected linked list to be empty")
		}
	})
	t.Run("Linked list has many data", func(t *testing.T) {
		nodeData := "head element"
		l.Prepend(&Node{Data: nodeData})
		l.Prepend(&Node{Data: nodeData})
		n := &Node{
			Data: nodeData,
		}
		l.Prepend(n)
		node, err := l.Shift()
		if err != nil {
			t.Error("Error returned, but not expected")
		}
		if node != n {
			t.Errorf("Returned %s but %s expected", node.Data, n.Data)
		}
		if l.head == nil {
			t.Error("Expected linked list not to be empty")
		}
	})
}

func TestLinkedList_Search(t *testing.T) {
	var l LinkedList
	t.Run("Linked list is empty", func(t *testing.T) {
		nodeData := "head element"
		result := l.Search(nodeData)
		if result {
			t.Error("Expected false but returned true")
		}
	})
	t.Run("Element exist and head", func(t *testing.T) {
		nodeData := "head element"
		newNode := &Node{
			Data: nodeData,
		}
		l.Append(newNode)
		result := l.Search(nodeData)
		if !result {
			t.Error("Expected true but returned false")
		}
	})
	t.Run("Element exist and not a head", func(t *testing.T) {
		nodeData := "second element"
		newNode := &Node{
			Data: nodeData,
		}
		l.Append(newNode)
		result := l.Search(nodeData)
		if !result {
			t.Error("Expected true but returned false")
		}
	})
	t.Run("Element does not exist", func(t *testing.T) {
		nodeData := "does not existing element"
		result := l.Search(nodeData)
		if result {
			t.Error("Expected false but returned true")
		}
	})
}

func TestLinkedList_Delete(t *testing.T) {
	var l LinkedList
	headData := "head data"
	secondData := "second data"
	notExistData := "not exist"
	newNode := &Node{
		Data: headData,
	}
	l.Append(newNode)
	newNode = &Node{
		Data: secondData,
	}
	l.Append(newNode)
	t.Run("delete not existing node", func(t *testing.T) {
		l.Delete(notExistData)
		if l.Search(notExistData) {
			t.Error("Expected node is not exist")
		}
	})
	t.Run("delete not a head node", func(t *testing.T) {
		l.Delete(secondData)
		if l.Search(secondData) {
			t.Errorf("Expected %s deleted", secondData)
		}
	})
	t.Run("delete head node", func(t *testing.T) {
		l.Delete(headData)
		if l.head != nil && l.head.Data == headData {
			t.Errorf("Expected node data not a %s", headData)
		}
	})
	t.Run("delete in empty list", func(t *testing.T) {
		l.Delete(notExistData)
		if l.Search(notExistData) {
			t.Error("Expected node is not exist")
		}
	})
}
