package queue

import (
	"bytes"
	"fmt"
)

type Queue struct {
	data        []interface{}
	tail, front int
}

func Instance() *Queue {
	return &Queue{
		data:  make([]interface{}, 6),
		tail:  0,
		front: 0,
	}
}

func (q *Queue) Enqueue(item interface{}) {
	// 首先先判断是否满，满了就扩容
	if q.IsFull() {
		q.resize(q.GetCapacity() * 2)
	}
	q.data[q.tail] = item
	q.tail = (q.tail + 1) % q.GetLen()
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}
	item := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.GetLen()
	return item
}

func (q *Queue) IsFull() bool {
	if (q.tail+1)%q.GetLen() == q.front {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if q.tail == q.front {
		return true
	}
	return false
}

func (q *Queue) GetSize() int {
	// 这里不用size来记录队列的长度，通过front和tail的位置来进行运算。
	// 因为是循环队列，如果front在tail后面，那直接相减，否则用容量减去空的数量得到size
	var size int
	if q.tail == q.front {
		return 0
	}

	if q.tail > q.front {
		size = q.tail - q.front
	} else {
		size = q.GetCapacity() - (q.front - q.tail)
	}
	return size
}

func (q *Queue) GetCapacity() int {
	return len(q.data) - 1
}

func (q *Queue) GetLen() int {
	return len(q.data)
}

func (q *Queue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity+1)
	for i := 0; i < q.GetSize(); i++ {
		index := (i + q.front) % q.GetLen()
		newData[i] = q.data[index]
	}
	//这里有个顺序问题，跟带size的不一样
	q.tail = q.GetSize()
	q.data = newData
	q.front = 0
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", q.GetSize(), q.GetCapacity()))
	buffer.WriteString("front [")
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(q.data[i]))
		if (i+1)%len(q.data) != q.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
