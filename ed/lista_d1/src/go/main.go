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
}

type LList struct {
	root *Node
}

func NewLList() *LList {
	root := &Node{}

	root.next = root
	root.prev = root

	return &LList{root: root}
}

func (ll *LList) size() int {
	node := ll.Front()
	count := 0
	for node != ll.root {
		count++
		node = node.next
	}
	return count
}

func (ll *LList) insert(prev, next, node *Node) {
	node.prev = prev
	node.next = next

	prev.next = node
	next.prev = node
}

func (ll *LList) remove(node *Node) {
	if node == ll.root {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (ll *LList) clear() {
	ll.root.prev = ll.root
	ll.root.next = ll.root
}

func (ll *LList) Front() *Node {
	return ll.root.next
}

func (ll *LList) Back() *Node {
	return ll.root.prev
}

func (ll *LList) Root() *Node {
	return ll.root
}

func (ll *LList) Size() int {
	return ll.size()
}

func (ll *LList) Clear() {
	ll.clear()
}

func (ll *LList) PushBack(data int) {
	node := &Node{data: data}
	ll.insert(ll.Back(), ll.Root(), node)
}

func (ll *LList) PushFront(data int) {
	node := &Node{data: data}
	ll.insert(ll.Root(), ll.Front(), node)
}

func (ll *LList) PopBack() {
	ll.remove(ll.Back())
}

func (ll *LList) PopFront() {
	ll.remove(ll.Front())
}

func (ll *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	node := ll.Front()
	for node != ll.Root() {
		if node.prev != ll.Root() {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", node.data))

		node = node.next
	}
	sb.WriteString("]")
	return sb.String()
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
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
