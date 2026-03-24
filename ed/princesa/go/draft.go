package main

import "fmt"
func main() {
	n, sword := 0, 0
	fmt.Scan(&n, &sword)
	sword--
	vet := make([]int, n)
	for i := 0; i < n; i++ {
		vet[i] = i+1
	}
	//fmt.Print(n, sword, vet)
	//show(vet, sword)
	alives := n 
	for alives > 1 {
		show(vet, sword)

		next_alive := index_next_vivo(vet, n, sword)
		vet[next_alive] = 0
		alives--
		
		sword = index_next_vivo(vet, n, sword)
	}
	show(vet, sword)
}

func index_next_vivo(vet []int, n int, sword int) int {
	next := (sword + 1) % n
	for vet[next] == 0 {
		next = (next+1) % n
	}
	return next
}

func show(vet []int, sword int) {
	fmt.Print("[ ")
	for i, v := range vet {
		if v != 0 {
			fmt.Print(v)
			if i == sword {
				fmt.Print(">")
			}
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
} 
