package arraystack

import "../array"

type Stack struct {
	array *array.Array
}

func (s *Stack) Instance() *Stack {
	return &Stack{array: array.Instance()}
}

func (s *Stack) Pop() interface{} {
	return s.array.RemoveLast()
}

func (s *Stack) Push(item interface{}) {
	s.array.AddLast(item)
}

func (s *Stack) Peek() interface{} {
	return s.array.Get(s.GetSize() - 1)
}

func (s *Stack) GetSize() int {
	return s.array.GetSize()
}

func (s *Stack) GetCapacity() int {
	return s.array.GetCapacity()
}

func (s *Stack) IsEmpty() bool {
	return s.array.IsEmpty()
}
