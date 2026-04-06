package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumngb(vet []int) []int {
	if len(vet) <= 1{
		return make([]int, 0)
	}
	sum := vet[0]+vet[1]
	return append([]int{sum}, sumngb(vet[1:])...)
}

func processa(vet []int) {
	// 1. defina o ponto de parada
	if len(vet) == 1 {
		fmt.Printf("[ %s ]\n", Join(vet, " "))
		return
	}
	// 2. monte o vetor auxiliar com os resultados das somas
	vetb := sumngb(vet)
	// 3. chame recursivamente a função processa para o vetor auxiliar
	processa(vetb)
	// 4. imprima o vetor original
	fmt.Printf("[ %s ]\n", Join(vet, " "))
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
