package stack

import (
	"github.com/kickbuttowskiii/go-algorithms/linked_list"
)

type Stack struct {
	Store *linked_list.LinkedList
}

func NewStack() *Stack {
	return &Stack{
		Store: &linked_list.LinkedList{},
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Store.Head == nil
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.Store.Head.Value
}

func (s *Stack) String() string {
	return s.Store.String()
}

func (s *Stack) Push(value int) {
	s.Store.Prepend(value)
}

func (s *Stack) Pop() int {
	node := s.Store.DeleteHead()
	if node == nil {
		return -1
	}
	return node.Value
}
