package queue

import "../array"

type Queue struct {
	Array *array.Array
}

func Instance() *Queue {
	return &Queue{
		Array: array.Instance(),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Array.AddLast(item)
}

func (q *Queue) Dequeue() interface{} {
	return q.Array.RemoveFirst()
}

func (q *Queue) GetFront() interface{} {
	return q.Array.GetFirst()
}

func (q *Queue) GetSize() int {
	return q.Array.GetSize()
}

func (q *Queue) GetCapacity() int {
	return q.Array.GetCapacity()
}

func (q *Queue) IsEmpty() bool {
	return q.Array.IsEmpty()
}
