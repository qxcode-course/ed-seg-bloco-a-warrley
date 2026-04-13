package main

import (
	"fmt"
)
func main() {
	var n, m int
	fmt.Scan(&n)
	vetn := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetn[i])
	}

	fmt.Scan(&m)
	vetm := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		a := 0
		fmt.Scan(&a)
		vetm[a] = true
	}

	// res := make([]int, m-n, 0)
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		value := vetn[i]
		if vetm[value] == false {
			res = append(res, value)
		} 
	}

	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
	fmt.Println()
}
