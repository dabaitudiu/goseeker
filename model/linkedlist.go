package model

import "github.com/pkg/errors"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	preHead *Node
	tail    *Node
	size    int
}

func NewNode(val interface{}) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

func NewLinkedList() *LinkedList {
	preHead := NewNode(0)

	return &LinkedList{
		preHead: preHead,
		tail:    preHead,
		size:    0,
	}
}

func (l *LinkedList) Add(val interface{}) {
	next := NewNode(val)
	if l.size == 0 {
		l.preHead.next = next
	}
	l.tail.next = next
	l.tail = next
	l.size += 1
}

func (l *LinkedList) AddFirst(val interface{}) {
	first := NewNode(val)
	first.next = l.preHead.next
	l.preHead.next = first
	l.size += 1
}

func (l *LinkedList) Find(val interface{}) bool {
	node := l.preHead
	for node.next != nil && node.next.val != val {
		node = node.next
	}
	return node.next != nil
}

func (l *LinkedList) Poll(val interface{}) error {
	if l.size == 0 {
		return errors.New("no nodes to drop")
	}
	curr := l.preHead
	// TODO: 注意这里直接比较会有一个坑点，如果interface的值是slice类型，直接比较会panic
	for curr != l.tail && curr.next.val != val {
		curr = curr.next
	}
	if curr == l.tail {
		return errors.New("no matched value found in list")
	}
	if curr.next == l.tail {
		l.tail = curr
	}
	curr.next = curr.next.next
	return nil
}
