package main

import "fmt"

func main() {
	var qtalbum, qtbaruel int
	fmt.Scan(&qtalbum, &qtbaruel)
	stickers := make([]int, qtbaruel)
	album := make([]int, qtalbum)

	for i := 0; i < qtbaruel; i++ {
		fmt.Scan(&stickers[i])
	}

	for i := 0; i < qtbaruel; i++ {
		album[stickers[i]-1]++
	}

	var repeated string
	var missing string

	for i := 0; i < qtalbum; i++ {
		if album[i] == 0 {
			missing += fmt.Sprintf("%d ", i+1)
		}
		if album[i] >= 1 {
			for j := 1; j < album[i]; j++ {
				repeated += fmt.Sprintf("%d ", i+1)
			}
		}
	}

	lrep := len(repeated)
	lmis := len(missing)
	if lrep > 0 {
		fmt.Printf("%s\n", repeated[:lrep-1])
	} else {
		fmt.Printf("%s\n", "N")
	}

	if lmis > 0 {
		fmt.Printf("%s\n", missing[:lmis-1])
	} else {
		fmt.Printf("%s\n", "N")
	}
}
