package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Perimeterizable interface {
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func describe(s Shape) {
	fmt.Printf("Area shape: %.2f\n", s.Area())
	if p, ok := s.(Perimeterizable); ok {
		fmt.Printf("perimeter: %.2f\n", p.Perimeter())
	}
}

func totalArea(shapes []Shape) float64 {
	total := float64(0)
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	shapes := []Shape{Circle{5.0}, Rectangle{5.0, 3.0}, Triangle{5.0, 3.2}}
	// fmt.Printf("%.2f - total area\n", totalArea(shapes))
	for _, shape := range shapes {
		describe(shape)
	}
}
