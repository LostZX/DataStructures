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
	l.AddLast(1)
	l.AddLast(5)
	l.AddFirst(3)
	fmt.Println(l)
	l = l.Remove(1)
	fmt.Println(l)
	l = l.Remove(5)
	fmt.Println(l)
}
