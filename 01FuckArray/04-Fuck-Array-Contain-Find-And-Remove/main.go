package main

import (
	"./array"
	"fmt"
)

func main() {
	a := array.Instance(5)
	a.Add(0, 1)
	a.Add(1, 2)
	a.Add(2, 3)
	fmt.Println(a)
	flag := a.Contains(1)
	fmt.Println(flag)
	index := a.Find(4)
	fmt.Println(index)
	a.Remove(1)
	fmt.Println(a)
	a.RemoveItem(1)
	fmt.Println(a)
	a.Add(1, 3)
	a.Add(2, 3)
	fmt.Println(a)
	a.RemoveAll()
	fmt.Println(a)
	fmt.Println(a.GetSize())
}
