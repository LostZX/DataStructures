package queue

import (
	"bytes"
	"fmt"
)

type Queue struct {
	data              []interface{}
	tail, front, size int
}

func Instance() *Queue {
	return &Queue{
		data:  make([]interface{}, 5),
		tail:  0,
		front: 0,
		size:  0,
	}
}

func (q *Queue) AddLast(item interface{}) {
	// 首先先判断是否满，满了就扩容
	if q.IsFull() {
		q.resize(q.GetCapacity() * 2)
	}
	q.data[q.tail] = item
	q.tail = (q.tail + 1) % q.GetCapacity()
	q.size++
}

func (q *Queue) AddFront(item interface{}) {
	// 首先先判断是否满，满了就扩容
	if q.IsFull() {
		q.resize(q.GetCapacity() * 2)
	}
	if q.IsEmpty() {
		q.AddLast(item)
		return
	}
	// 加到前前面去，防止 -1 下标越界 取模的值，必定会模的范围内；所以，计算机领域引用该特性，使元素路由算法不超出边界，并有规则存放。
	q.front = (q.front - 1 + q.GetCapacity()) % q.GetCapacity()
	q.data[q.front] = item
	q.size++
}

func (q *Queue) RemoveFront() interface{} {
	ref := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.GetCapacity()
	q.size--
	return ref
}

func (q *Queue) RemoveLast() interface{} {
	ref := q.data[q.tail-1]
	q.data[q.tail-1] = nil
	q.tail = (q.tail - 1 + q.GetCapacity()) % q.GetCapacity()
	q.size--
	return ref
}

func (q *Queue) IsFull() bool {
	if q.size == q.GetCapacity() {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue) GetSize() int {
	// 因为复制的上一个文件夹，有些地方不想改了
	return q.size
}

func (q *Queue) GetCapacity() int {
	// 不预留空位，无需-1
	return len(q.data)
}
func (q *Queue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < q.GetSize(); i++ {
		index := (i + q.front) % q.GetCapacity()
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
