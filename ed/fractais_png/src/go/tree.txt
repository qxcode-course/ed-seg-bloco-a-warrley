package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func drawTree(p *Pen, size float64) {
	if size < 2 {
		return
	}
	
	// p.SavePNG("tree.png")
	// var trash rune
	// fmt.Scanf("%c", &trash)
	
	degree := 35.0

	p.Walk(size)
	p.Right(degree)				
	drawTree(p, size*0.66)

	p.Left(2*degree)
	drawTree(p, size*0.66)

	p.Right(degree)
	p.Walk(-size)
}

func main() {
	pen := NewPen(1000, 1000)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(500, 1000) // move o pincel para x 250, y 500
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.SetLineWidth(1)

	drawTree(pen, 350)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}