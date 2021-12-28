package main

import (
	"./array"
	"./student"
	"fmt"
)

func main() {
	a := array.Instance(2)
	a.Add(0, student.Instance("zhangsan", 0))
	a.Add(1, student.Instance("lisi", 6))
	fmt.Println(a.Get(1))

}
