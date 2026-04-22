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

func (s *set) InsertOrder(value int) {
	exist, id := s.BinarySearch(value)
	if exist {
		return
	}

	if s.size == s.capacity {
		s.Reserve(s.capacity * 2)
	}

	for i := s.size; i > id; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[id] = value
	s.size++
}

func (s *set) Erase(value int) {
	exist, id := s.BinarySearch(value)
	if !exist {
		fmt.Println("value not found")
		return
	}

	for i := id; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
}

func (s *set) BinarySearch(value int) (bool, int) {
	low := 0
	high := s.size - 1

	for low <= high {
		mid := (low + high) / 2
		if s.data[mid] == value {
			return true, mid
		} else if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false, low
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)

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
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.InsertOrder(value)
			}
		case "show":
			fmt.Println(s)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			s.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			exist, _ := s.BinarySearch(value)
			fmt.Println(exist)
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
