package main

import (
	"./stack"
	"fmt"
)

func main() {
	s := stack.Instance()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	top := s.Peek()
	fmt.Println(top)
	fmt.Println(s)
	p := s.Pop()
	fmt.Println(p)
	fmt.Println(s)
}
