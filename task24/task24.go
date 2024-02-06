// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func length(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func Run() {
	var x, y float64
	fmt.Println("Задайте координаты первой точки")
	fmt.Scan(&x, &y)
	p1 := Point{x, y}
	fmt.Println("Задайте координаты второй точки")
	fmt.Scan(&x, &y)
	p2 := Point{x, y}
	fmt.Println("Расстояние между точками -", length(p1, p2))
}
