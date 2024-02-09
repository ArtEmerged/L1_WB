package main

import (
	"fmt"
	"wblone/tasks/24/point"
)

func main() {
	p1 := point.NewPoint(10.4, 2.4)
	p2 := point.NewPoint(1.1, 8.4)
	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками a%.2f и b%.2f = %.3f\n", p1, p2, distance)

}
