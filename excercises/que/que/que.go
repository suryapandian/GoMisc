package que

import (
	"fmt"
	"que.com/stack"
)

type Queue struct {
	stack.Stack
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Add(val int) {
	fmt.Println("Add to queue")
	q.Push(val)

}

func (q *Queue) Delete() {
	fmt.Println("Delete from queue")
	q.Pop()
}
