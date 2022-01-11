package queue

import (
	"bytes"
	"fmt"
)

type Queue struct {
	data              []interface{}
	front, tail, size int
}

func Instance() *Queue {
	return &Queue{
		data:  make([]interface{}, 6),
		size:  0,
		tail:  0,
		front: 0,
	}
}

func (q *Queue) IsFull() bool {
	// 求模的意义
	if (q.tail+1)%len(q.data) == q.front {
		return true
	}
	return false
}

func (q *Queue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	return q.data[q.front]
}

func (q *Queue) GetLen() int {
	return len(q.data)
}

func (q *Queue) Enqueue(item interface{}) {
	if q.IsFull() {
		q.resize(q.GetCapacity() * 2)
	}

	q.data[q.tail] = item
	q.size++
	// 循环队列， 这里不用GetCapacity是因为要算上那个空格
	q.tail = (q.tail + 1) % q.GetLen()
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	item := q.data[q.front]
	// 由于是循环队列，所以置空不删除
	q.data[q.front] = nil
	// 去模数 计算队首的游标
	q.front = (q.front + 1) % q.GetCapacity()
	q.size--

	return item
}

func (q *Queue) GetCapacity() int {
	// -1 与初始化+1保持一致
	return len(q.data) - 1
}

func (q *Queue) IsEmpty() bool {
	return q.front+1 == q.tail
}

func (q *Queue) resize(capacity int) {
	// +1 的意义，保持始终有一个空位，-> 判断 IsEmpty的方法 + 1
	newData := make([]interface{}, capacity+1)
	for i := 0; i < q.size; i++ {
		// debug
		index := (q.front + i) % q.GetLen()
		newData[i] = q.data[index]
	}
	q.data = newData
	// 扩容后 tail就是原来的size，front就是0，
	q.front = 0
	q.tail = q.size
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", q.size, q.GetCapacity()))
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
