package main

import "fmt"

type Rectangle struct {
	length, width float64
}

func Area(r *Rectangle) float64 {
	return r.length * r.width
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func Scale(r *Rectangle, f float64) {
	r.length *= f
	r.width *= f
}

func (r *Rectangle) Scale(f float64) {
	r.length *= f
	r.width *= f
}

func main() {
	r := &Rectangle{1, 2}
	fmt.Println(Area(r))
	Scale(r, 10)
	r.Scale(10)
	fmt.Println(*r)
	fmt.Println(r.Area())
}
