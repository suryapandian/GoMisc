package stack

import "fmt"

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Pop() {
	fmt.Println("Popping")
	stackSize := len(*s)
	if stackSize == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(val int) {
	fmt.Println("Pushing", val)
	*s = append(*s, val)
}
