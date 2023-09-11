package main

import (
	"fmt"
	"math"
)

// 24. Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

// Т.к. поля инкапсулированы мы не имеем к ним прямого доступа и пишем с маленькой буквы
type Point struct {
	x, y float64
}

// Конструктор создающий новую точку
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между точками
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	// находим разницу между точками
	betweenX := p1.x - p2.x
	betweenY := p1.y - p2.y
	// вычиляем результат используя формулу d = √ (x2-x1)² + (y2-y1)²
	return math.Sqrt(betweenX*betweenX + betweenY*betweenY)
}

func main() {
	// Создание двух точек.
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисление расстояния между точками
	distance := point1.DistanceTo(point2)

	// Вывод результата. Используем %.2f для вывода первых 2 нулей в дробной части
	fmt.Printf("Расстояние между точкой 1 (%.2f, %.2f) и точкой 2 (%.2f, %.2f): %.2f\n", point1.x, point1.y, point2.x, point2.y, distance)
}
