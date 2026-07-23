package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	area() (float64, error)
	perimeter() (float64, error)
}

func viewShape(s Shape) {
	fmt.Printf("%-18T| ", s)

	area, errA := s.area()
	if errA != nil {
		fmt.Println(errA)
		return
	}
	perimeter, errP := s.perimeter()
	if errP != nil {
		fmt.Println(errP)
		return
	}

	fmt.Printf("Area - %6.2f | Perimeter - %6.2f \n", area, perimeter)
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0, errors.New("height and width must be greater than 0")
	}
	return r.width * r.height, nil
}
func (r Rectangle) perimeter() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0, errors.New("height and width must be greater than 0")
	}
	return 2*r.width + 2*r.height, nil
}

type Circle struct {
	radius float64
}

func (c Circle) area() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("radius must be greater than 0")
	}
	return math.Pi * c.radius * c.radius, nil
}
func (c Circle) perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("radius must be greater than 0")
	}
	return 2 * math.Pi * c.radius, nil
}

type Triangle struct {
	base, height, a, b float64
}

func (t Triangle) area() (float64, error) {
	if t.base <= 0 || t.height <= 0 {
		return 0, errors.New("base and height must be greater than 0")
	}
	return 0.5 * t.base * t.height, nil
}
func (t Triangle) perimeter() (float64, error) {
	if t.base <= 0 || t.height <= 0 {
		return 0, errors.New("base and height must be greater than 0")
	}
	return t.base + t.a + t.b, nil
}

type Parallelogram struct {
	base, height, side float64
}

func (p Parallelogram) area() (float64, error) {
	if p.base <= 0 || p.height <= 0 {
		return 0, errors.New("base and height must be greater than 0")
	}
	return p.base * p.height, nil
}
func (p Parallelogram) perimeter() (float64, error) {
	if p.base <= 0 || p.height <= 0 {
		return 0, errors.New("base and height must be greater than 0")
	}
	return 2*p.base + 2*p.side, nil
}

func main() {
	shapes := []Shape{
		Rectangle{10, 20},
		Circle{10},
		Triangle{6, 4, 5, 3},
		Parallelogram{6, 0, 2},
	}

	for _, shape := range shapes {
		viewShape(shape)
	}
}
