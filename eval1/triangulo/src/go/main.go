package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printvet(vet []int) {
	fmt.Print("[ ")
	for i := 0; i < len(vet); i++ {
		fmt.Printf("%d ", vet[i])
	}
	fmt.Println("]")
}

func sumNeig (vet []int, i int) []int{
	if len(vet)-1 == i {
		return []int{}
	}

	sum := vet[i]+vet[i+1]
	return append([]int{sum}, sumNeig(vet, i+1)...)
}

func processa(vet []int) {
	if len(vet) == 1 {
		printvet(vet)
		return
	}

	vecb := sumNeig(vet, 0)
	processa(vecb)
	printvet(vet)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
