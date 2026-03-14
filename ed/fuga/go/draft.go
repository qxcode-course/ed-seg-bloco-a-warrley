package main

import "fmt"

func mod16 (x int) int {
	return (x+16) % 16
}

func main() {
	var h, p, f, d int
	fmt.Scanf("%d %d %d %d", &h, &p, &f, &d)

	var dp, dh int
	if d == -1 {
		dh = mod16(f-h)
		dp = mod16(f-p)
	} else {
		dh = mod16(h - f)
		dp = mod16(p - f)

	}
	var res string

	if dp < dh {
		res = "N"
	} else {
		res = "S"
	}

	fmt.Printf("%s\n", res)
}
