package main

import (
	"fmt"
	"wblone/tasks/24/point"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуДары Point с инкапсулированными параметрами x,y и конструктором.
*/

func main() {
	// Создаем точки p1, p2
	p1 := point.NewPoint(1, 1)
	p2 := point.NewPoint(100, 1)

	// Вычисляем расстояние между точками p1 и p2
	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками a%.2f и b%.2f = %.3f\n", p1, p2, distance)

}
