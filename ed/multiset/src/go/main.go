package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type multiset struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *multiset {
	var ms multiset
	ms.data = make([]int, capacity)
	ms.capacity = capacity
	ms.size = 0

	return &ms
}

func (ms *multiset) expand() {
	ndata := make([]int, 2*ms.capacity)

	for i := range ms.size {
		ndata[i] = ms.data[i]
	}

	ms.data = ndata
	ms.capacity = 2 * ms.capacity
}

func (ms *multiset) search(value int) (bool, int) {
	low, high := 0, ms.size

	for low < high {
		mid := (low + high) / 2

		if ms.data[mid] <= value {
			low = mid + 1
		} else {
			high = mid
		}
	}

	pos := low - 1

	if pos >= 0 && ms.data[pos] == value {
		return true, pos
	}

	return false, low
}

func (ms *multiset) insert(value, index int) error {
	if ms.size == ms.capacity || index == ms.capacity {
		ms.expand()
	}

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++

	return nil
}

func (ms *multiset) Insert(value int) {
	// fmt.Printf("\nsize:%d\n capacity:%d\n value:%d\n", ms.size, ms.capacity, value)
	exist, index := ms.search(value)

	if exist {
		ms.insert(value, index+1)
	} else {
		ms.insert(value, index)
	}
}

func (ms *multiset) Contains(value int) bool {
	exist, _ := ms.search(value)
	return exist
}

func (ms *multiset) erase(index int) error {
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
	return nil
}

func (ms *multiset) Erase(value int) error {
	exist, index := ms.search(value)

	if !exist {
		return fmt.Errorf("value not found\n")
	} else {
		ms.erase(index)
	}

	return nil
}

func (ms *multiset) Count(value int) int {
	// fmt.Printf("index=%d", index)
	_, index := ms.search(value)
	count := 0
	for i := index; ms.size > 0 && i >= 0 && ms.data[i] == value; i-- {
		count++
	}

	return count
}

func (ms *multiset) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}

	return count
}

func (ms *multiset) Clear() {
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *multiset) String() string {
	return fmt.Sprintf("[%s]", Join(ms.data[0:ms.size], ", "))
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ok := ms.Erase(value)
			if ok != nil {
				fmt.Print(ok)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			// ms.Clear()
			ms.size = 0
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
