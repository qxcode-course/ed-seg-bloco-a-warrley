package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data int
	prev *Node
	next *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root

	ll := &LList{root: root}
	return ll
}

func (ll *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	curr := ll.Front()
	for curr != nil {
		sb.WriteString(fmt.Sprintf("%d", curr.data))

		if curr.Next() != nil {
			sb.WriteString(", ")
		}
		curr = curr.Next()
	}
	sb.WriteString("]")
	return sb.String()
}

func (ll *LList) Front() *Node {
	return ll.root.Next()
}

func (ll *LList) Back() *Node {
	return ll.root.Prev()
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Insert(node *Node, value int) {
	prev := node.prev

	n := &Node{
		data: value,
		prev: prev,
		next: node,
		root: ll.root,
	}

	prev.next = n
	node.prev = n

	ll.size++
}

func (ll *LList) PushBack(value int) {
	ll.Insert(ll.root, value)
}

func (ll *LList) PushFront(value int) {
	ll.Insert(ll.root.next, value)
}

func (ll *LList) Remove(node *Node) *Node {
	next := node.Next()

	node.prev.next = node.next
	node.next.prev = node.prev

	ll.size--
	return next
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) Search(value int) *Node {
	for curr := ll.Front(); curr != nil; curr = curr.Next() {
		if curr.data == value {
			return curr
		}
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll)
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.data)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.data)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.data = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
