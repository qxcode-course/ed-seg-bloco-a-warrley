package main

import "fmt"
func main() {
    n := 0
    fmt.Scan(&n)

    vet := make([]int, n)
    singles := make(map[int]int)

    count := 0
    for i := 0; i < n; i++ {
        value := 0
        fmt.Scan(&value)
        if singles[-value] % 2 != 0 {
            singles[-value]--
            count++
        } else {
            singles[value]++
        }
    }
    
    for i := 0; i < n; i++ {
        value := vet[i]
        if singles[value] == 2 {
            count++
        }
    }

    fmt.Println(count)
}
