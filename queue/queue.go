package queue

import (
	"github.com/kickbuttowskiii/go-algorithms/linked_list"
)

type Queue struct {
	Store *linked_list.LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		Store: &linked_list.LinkedList{},
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Store.Head == nil
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.Store.Head.Value
}

func (q *Queue) Enqueue(value int) {
	q.Store.Append(value)
}

func (q *Queue) Dequeue() int {
	node := q.Store.DeleteHead()

	if node == nil {
		return -1
	}

	return node.Value
}

func (q *Queue) String() string {
	return q.Store.String()
}
