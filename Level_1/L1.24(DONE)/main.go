package main

import (
	"fmt"
	"l_point/point"
)

func main() {

	// создаем две точки
	p1 := point.CreatePoint(1, 7)
	p2 := point.CreatePoint(4, 3)

	distance := p1.GetDistance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
