package linked_list

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.Value)
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *LinkedList) Prepend(value int) *LinkedList {
	node := &Node{
		Value: value,
		Next:  ll.Head,
	}

	ll.Head = node

	if ll.Tail == nil {
		ll.Tail = node
	}

	return ll
}

func (ll *LinkedList) Array() []int {
	var nodes []int
	currentNode := ll.Head

	for currentNode != nil {
		nodes = append(nodes, currentNode.Value)
		currentNode = currentNode.Next
	}

	return nodes
}

func (ll *LinkedList) String() string {
	var result string
	nodes := ll.Array()

	for _, node := range nodes {
		result += fmt.Sprintf("%d ", node)
	}

	return result
}

func (ll *LinkedList) Append(value int) *LinkedList {
	node := &Node{
		Value: value,
	}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return ll
	}

	ll.Tail.Next = node
	ll.Tail = node
	return ll
}

func (ll *LinkedList) FromArray(values []int) {
	for _, value := range values {
		ll.Append(value)
	}
}

func (ll *LinkedList) Insert(value, idx int) *LinkedList {
	if idx == 0 {
		ll.Prepend(value)
		return ll
	}

	count := 1
	node := &Node{Value: value}
	currentNode := ll.Head

	for currentNode != nil {
		if idx == count {
			break
		}
		currentNode = currentNode.Next
		count++
	}

	if currentNode != nil {
		node.Next = currentNode.Next
		currentNode.Next = node
	} else {
		if ll.Tail != nil {
			ll.Tail.Next = node
			ll.Tail = node
		} else {
			ll.Head = node
			ll.Tail = node
		}
	}

	return ll
}

func (ll *LinkedList) Find(value int) *Node {
	if ll.Head == nil {
		return nil
	}

	currentNode := ll.Head
	for currentNode != nil {
		if currentNode.Value == value {
			break
		}
		currentNode = currentNode.Next
	}

	return currentNode
}

func (ll *LinkedList) DeleteHead() *Node {
	if ll.Head == nil {
		return nil
	}

	headNode := ll.Head
	if ll.Head.Next != nil {
		ll.Head = ll.Head.Next
	} else {
		ll.Head = nil
		ll.Tail = nil
	}

	return headNode
}

func (ll *LinkedList) DeleteTail() *Node {
	node := ll.Tail

	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
		return nil
	}

	currentNode := ll.Head
	for currentNode.Next != nil {
		if currentNode.Next.Next == nil {
			currentNode.Next = nil
		} else {
			currentNode = currentNode.Next
		}
	}

	ll.Tail = currentNode
	return node
}
