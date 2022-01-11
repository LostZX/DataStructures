package main

import (
	"./queue"
	"fmt"
)

func main() {
	queue := queue.Instance()
	for i := 0; i < 10; i++ {
		queue.AddLast(i)
		fmt.Println(queue)
	}
	queue.AddFront(10)
	fmt.Println(queue)
	queue.RemoveLast()
	fmt.Println(queue)
	queue.RemoveFront()
}
