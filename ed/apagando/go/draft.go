package main

import "fmt"
func main() {
	n, m := 0, 0
	fmt.Scan(&n)
	vet := make([]int, n)
	for i := range n {
		fmt.Scan(&vet[i])
	}
	fmt.Scan(&m)
	del := make(map[int]struct{}, 0)
	for i := 0; i < m; i++ {
		a := 0
		fmt.Scan(&a)
		del[a] = struct{}{}	
	}

	res := make([]int, 0, n)
	for _, v := range vet {
		if _, ok := del[v]; !ok{
			res = append(res, v)
		}
	}

	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}