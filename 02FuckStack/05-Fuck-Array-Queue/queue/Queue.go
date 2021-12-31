package queue

import "../array"

type Queue struct {
	array *array.Array
}

func Instance() *Queue {
	return &Queue{
		array: array.Instance(),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.array.AddLast(item)
}

func (q *Queue) Dequeue() interface{} {
	return q.array.RemoveFirst()
}

func (q *Queue) GetFront() interface{} {
	return q.array.GetFirst()
}

func (q *Queue) GetSize() int {
	return q.array.GetSize()
}

func (q *Queue) GetCapacity() int {
	return q.array.GetCapacity()
}

func (q *Queue) IsEmpty() bool {
	return q.array.IsEmpty()
}
