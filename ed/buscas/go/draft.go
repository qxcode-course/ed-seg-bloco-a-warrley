package main

import "fmt"

func matchingStrings (vcon []string, vbus []string) []int {
    res := make([]int, len(vbus))
    freq := make(map[string]int)

    for _, str := range vcon {
        freq[str]++
    }

    for i, str := range vbus {
        res[i] = freq[str]
    }

    return res
}

func main() {
    var ncon, nbus int
    fmt.Scan(&ncon)
    vcon := make([]string, ncon)
    for i := 0; i < ncon; i++ {
        fmt.Scan(&vcon[i])
    }
    
    fmt.Scan(&nbus)
    vbus := make([]string, nbus)
    for i := 0; i < nbus; i++ {
        fmt.Scan(&vbus[i])
    }

    res := matchingStrings(vcon, vbus)
    for i, res := range res {
        if i != 0 {
            fmt.Print(" ")
        }
        fmt.Print(res)
    }
    fmt.Println()
}