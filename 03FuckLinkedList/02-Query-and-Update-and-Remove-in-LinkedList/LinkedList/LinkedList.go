package LinkedList

import (
	"bytes"
	"fmt"
)

type LinkedList struct {
	size      int
	dummyHead *Node // 虚拟节点，不计入size
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type Node struct {
	e    interface{}
	next *Node
}

func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}

func Instance() *LinkedList {
	return &LinkedList{dummyHead: &Node{}}
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Set(index int, e interface{}) {
	node := l.GetPoint(index)
	node.e = e
}

func (l *LinkedList) Get(index int) interface{} {
	return l.GetPoint(index).e
}

func (l *LinkedList) GetPoint(index int) *Node {
	if index < 0 || index >= l.size {
		panic("get failed, index is out of range")
	}
	cur := l.dummyHead
	i := 0
	for i != index {
		cur = cur.next
		i++
	}

	return cur.next
}

// RemoveFirst 为了实现一个栈
func (l *LinkedList) RemoveFirst() interface{} {
	if l.size == 0 {
		panic("remove failed, index is out of range.")
	}

	node := l.dummyHead.next
	l.dummyHead.next = node.next
	l.size--
	return node.e
}

func (l *LinkedList) Remove(e interface{}) *LinkedList {
	cur := l.dummyHead
	for cur.next != nil {
		if cur.next.e == e {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	l.size--
	return l
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LinkedList) Add(index int, e interface{}) {
	if 0 > index || index > l.size {
		panic("Add failed, Illegal index.")
	}

	prev := l.dummyHead

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e, prev.next}
	l.size++
}

func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur.next != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}
