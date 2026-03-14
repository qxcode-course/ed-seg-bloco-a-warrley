package main

import "fmt"

type gome struct {
	x int
	y int
}

func main() {
	var q int
	var d string
	fmt.Scan(&q, &d)

	snake := make([]gome, q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &snake[i].x, &snake[i].y)
	}

	for i := q - 1; i > 0; i-- {
		snake[i] = snake[i-1]
	}

	if d == "U" {
		snake[0].y--
 	} else if d == "D" {
		snake[0].y++
 	} else if d == "R" {
		snake[0].x++
 	} else if d == "L" {
		snake[0].x--
 	}

	printSnake(snake, q)
}

func printSnake(snake []gome, qt int) {
	for i := 0; i < qt; i++ {
		fmt.Printf("%d %d\n", snake[i].x, snake[i].y)
	}
}