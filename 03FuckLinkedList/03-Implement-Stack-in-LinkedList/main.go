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
	s.Push(4)
	s.Push(5)
	s.Push(2)
	s.Push(10)
	fmt.Println(s)
}
