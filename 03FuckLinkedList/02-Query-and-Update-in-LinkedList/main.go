package main

import (
	"./LinkedList"
	"fmt"
)

func main() {
	l := LinkedList.Instance()
	l.AddLast(1)
	l.AddFirst(2)
	l.Set(1, 5)
	fmt.Println(l.Contains(6))
}
