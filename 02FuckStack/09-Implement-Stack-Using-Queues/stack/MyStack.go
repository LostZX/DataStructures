package stack

import (
	"../queue"
)

type MyStack struct {
	// 用队列模拟一个栈
	Queue *queue.Queue
	Tmp   *queue.Queue
}

func Constructor() MyStack {
	return MyStack{
		Queue: queue.Instance(),
		Tmp:   queue.Instance(),
	}
}

func (s *MyStack) Push(x int) {
	s.Queue.Enqueue(x)
	// 如果长度是 1 就结束
	if s.Queue.GetSize() == 1 {
		return
	}

	// 先取出size，不然出队会影响循环
	qSize := s.Queue.GetSize()

	for i := 0; i < qSize-1; i++ {
		ref := s.Queue.Dequeue()
		s.Tmp.Enqueue(ref)
	}

	// 先取出size，不然出队会影响循环
	tSize := s.Tmp.GetSize()

	for i := 0; i < tSize; i++ {
		ref := s.Tmp.Dequeue()
		s.Queue.Enqueue(ref)
	}
}

func (s *MyStack) Pop() int {
	return s.Queue.Dequeue().(int)
}

func (s *MyStack) Top() int {
	// 使用底层数组拿到栈顶, 其实应用两个队列便利一遍，但是懒得写了
	last := s.Queue.Array.GetLast()
	return last.(int)
}

func (s *MyStack) Empty() bool {
	return s.Queue.IsEmpty()
}
