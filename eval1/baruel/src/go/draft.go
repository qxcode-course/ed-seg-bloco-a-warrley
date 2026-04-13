package main

import "fmt"
func main() {
    album_total, quantidade := 0, 0
    fmt.Scanf("%d", &album_total)
    fmt.Scanf("%d", &quantidade)

    count := make(map[int]int)
    repeated := make([]int, 0)

    for i := 0; i < quantidade; i++ {
        value := 0
        fmt.Scan(&value)

        count[value]++

        if count[value] > 1 {
            repeated = append(repeated, value)
        }
    }

    if len(repeated) == 0 {
        fmt.Println("N")
    } else {
        for i, v := range repeated {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(v)
        }
        fmt.Println()
    }

    for i := 1; i <= album_total; i++ {
        if i != 1 {
            fmt.Print(" ")
        }
        if count[i] == 0 {
            fmt.Print(i)
        }
    }
    fmt.Println()

    // fmt.Print(album_total, quantidade, count, repeated)
}
