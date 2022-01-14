package stack

import (
	"../LinkedList"
	"bytes"
)

type Stack struct {
	list *LinkedList.LinkedList
	size int
}

func Instance() *Stack {
	return &Stack{LinkedList.Instance(), 0}
}

func (s *Stack) Push(e interface{}) {
	s.list.AddFirst(e)
	s.size++
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) Pop() interface{} {
	s.size--
	return s.list.RemoveFirst()
}

func (s *Stack) Peek() interface{} {
	return s.list.Get(0)
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Stack: top ")
	buffer.WriteString(s.list.String())

	return buffer.String()
}
