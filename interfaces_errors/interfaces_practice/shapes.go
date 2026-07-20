package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func viewShape(s Shape) {
	fmt.Printf("%T: Area - %f; Perimeter - %f \n", s, s.area(), s.perimeter())
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	viewShape(Rectangle{3, 4})
	viewShape(Circle{10})
}
