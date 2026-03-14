package main

import (
	"fmt"
	"math"
)
func main() {
    var qt int
    fmt.Scanf("%d", &qt)
    launches := make([]int, qt)
    for i := 0; i < qt; i++ {
        var l1, l2 int
        fmt.Scanf("%d %d", &l1, &l2)

        if l1 < 10 || l2 < 10 {
            launches[i] = -1
        } else {
            launches[i] = int(math.Abs(float64(l1-l2)))
        }
    }

    vwin := 101
    iwin := 101
    for i := 0; i < qt; i++ {
        if launches[i] != -1 && launches[i] < vwin {
            vwin = launches[i]
            iwin = i
        }
    }

    if iwin != 101 {
        fmt.Println(iwin)
    } else {
        fmt.Println("sem ganhador")
    }
}
