package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data     int
	hassword bool
	prev     *Node
	next     *Node
	root     *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return n.root.next
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return n.root.prev
	}
	return n.prev
}

type CList struct {
	root *Node
	size int
}

func NewCList() *CList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root

	cl := &CList{root: root, size: 0}
	return cl
}

func (cl *CList) Size() int {
	return cl.size
}

func (cl *CList) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for curr := cl.root.next; curr != curr.root; curr = curr.next {
		if curr.hassword {
			sb.WriteString(fmt.Sprintf("%d> ", curr.data))
		} else {
			sb.WriteString(fmt.Sprintf("%d ", curr.data))
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (cl *CList) Front() *Node {
	return cl.root.Next()
}

func (cl *CList) Back() *Node {
	return cl.root.Prev()
}

func (cl *CList) Insert(node *Node, value int) {
	prev := node.prev

	n := &Node{
		data:     value,
		hassword: false,
		prev:     prev,
		next:     node,
		root:     cl.root,
	}

	prev.next = n
	node.prev = n

	cl.size++
}

func (cl *CList) PushBack(value int) {
	cl.Insert(cl.root, value)
}

func (cl *CList) PushFront(value int) {
	cl.Insert(cl.root.next, value)
}

func (cl *CList) Remove(node *Node) *Node {
	next := node.Next()

	node.prev.next = node.next
	node.next.prev = node.prev

	cl.size--
	return next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewCList()
	var killer *Node
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
		if i == chosen {
			l.Back().hassword = true
			killer = l.Back()
		}
	}

	for l.Size() > 1 {
		fmt.Println(l)

		killer.hassword = false
		l.Remove(killer.Next())

		killer = killer.Next()
		killer.hassword = true
	}

	fmt.Println(l)
}
