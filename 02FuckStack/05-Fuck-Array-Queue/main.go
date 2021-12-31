package main

import (
	"./queue"
	"fmt"
)
import "../../student"

func main() {
	queue := queue.Instance()
	s1 := student.New("yi", 100)
	s2 := student.New("er", 200)
	s3 := student.New("san", 300)

	queue.Enqueue(s1)
	queue.Enqueue(s2)
	queue.Enqueue(s3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.GetSize())
}
