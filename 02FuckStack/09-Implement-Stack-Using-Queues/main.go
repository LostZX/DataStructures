package main

import (
	"./stack"
	"fmt"
)

func main() {
	s := stack.Constructor()
	for i := 0; i < 10; i++ {
		s.Push(i)
		fmt.Println(s.Queue)
	}

	ref := s.Pop()
	fmt.Println(ref)
	fmt.Println(s.Queue)
	fmt.Println(s.Top())
}
