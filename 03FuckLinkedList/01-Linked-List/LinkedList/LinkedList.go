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

func (l *LinkedList) AddFirst(e interface{}) {
	node := &Node{e, nil}
	if l.size == 0 {
		l.dummyHead.next = node
	} else {
		ref := l.dummyHead.next
		l.dummyHead.next = node
		l.dummyHead.next.next = ref
	}
	l.size++
}

func (l *LinkedList) AddLast(e interface{}) {
	node := &Node{e, nil}
	if l.size == 0 {
		l.dummyHead.next = node
	} else {
		cur := l.dummyHead
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	l.size++
}
