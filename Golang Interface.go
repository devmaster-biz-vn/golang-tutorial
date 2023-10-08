package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func totalArea(s []Shape) float64 {
	total := 0.0
	for _, shape := range s {
		total += shape.Area()
	}
	return total
}

func main() {
	c := Circle{radius: 35}
	var v Shape = c
	fmt.Println("Area of Circle: ", v.Area())

	rec := Rectangle{width: 5.34, height: 7.89}
	fmt.Println("Area of Rectangle: ", rec.Area())

	shape := []Shape{c, rec}
	fmt.Println("Total Area: ", totalArea(shape))
}
