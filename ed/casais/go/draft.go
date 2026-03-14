package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vet := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
	}

	var count int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// fmt.Println(vet)
			if vet[i] != 0 && vet[i]+vet[j] == 0 {
				// fmt.Printf("%d-%d\n", vet[i], vet[j])
				count++
				vet[i] = 0
				vet[j] = 0
			}
		}
	}

	fmt.Printf("%d\n", count)

}

// func abs(x int) int{
//     if x < 0 {
//         return -x
//     } else {
//         return x
//     }
// }
