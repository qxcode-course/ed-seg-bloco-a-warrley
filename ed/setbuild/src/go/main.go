package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type set struct {
	data     []int
	size     int
	capacity int
}

func (s *set) String() string {
	return fmt.Sprint("[" + Join(s.data[0:s.size], ", ") + "]")
}

func (s *set) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", s.size, s.capacity)
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func NewSet(capacity int) *set {
	return &set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *set) Reserve(newCapacity int) {
	newData := make([]int, newCapacity)

	// fmt.Println(newData.Status())
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	// s = &newsector NAO FUNCIONA
	s.data = newData
	s.capacity = newCapacity
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
			fmt.Println(v)
		case "erase":
			// value, _ := strconv.Atoi(parts[1])
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
