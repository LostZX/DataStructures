package main

import (
	"./array"
	"./student"
	"fmt"
)

func main() {
	a := array.Instance()
	a.Add(0, student.Instance("zhangsan", 0))
	a.Add(1, student.Instance("lisi", 6))
	fmt.Println(a.Get(1))
	a.Add(2, student.Instance("wangwu", 100))
	a.Add(2, student.Instance("wangwu", 100))
	a.Add(2, student.Instance("wangwu", 100))
	a.Add(2, student.Instance("wangwu", 100))
	a.Add(2, student.Instance("wangwu", 100))
	fmt.Println(a)
}
