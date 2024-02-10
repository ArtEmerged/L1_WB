package point

import "math"

// Создаем структуру Point с инкапсулированными параметрами x, y
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для определния расстояния между двумя точкми по формуле AB = √((xb - xa)^2 + (yb - ya)^2)
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}
