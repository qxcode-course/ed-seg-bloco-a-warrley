package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}


func (p *Pen) drawWeizen(size float64) {
	queue := []PenState{{p.x, p.y, p.angle, size}}

	for len(queue) > 0 {
		ht := queue[0]
		queue = queue[1:]

		p.x, p.y, p.angle = ht.x, ht.y, ht.angle

		if ht.size < 2 {
			continue
		}

		p.Walk(ht.size)

		for range(6) {
			
		}
	}

}

func main() {
	pen := NewPen(1000, 1000)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetPosition(500, 950) // move o pincel para x 250, y 500
	pen.SetLineWidth(1.5)
	pen.SetRGB(255, 255, 255)
	pen.SetHeading(90)

	pen.drawWeizen(300)

	pen.SavePNG("weizen.png")
	fmt.Println("PNG file created successfully.")
}
