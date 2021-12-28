package main

import (
	"./array"
	"fmt"
)

func main() {
	a := array.Instance(5)
	a.Add(0, 1)
	a.Add(1, 2)
	fmt.Println(a)
	a.AddFirst(3)
	fmt.Println(a)
	a.AddLast(5)
	fmt.Println(a)
}
