package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	Area() float64
}

// Define a struct
type Circle struct {
	Radius float64
}

// Implement the Area() method for the Circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Define another struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area() method for the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	s = c
	fmt.Println("Circle Area:", s.Area())

	s = r
	fmt.Println("Rectangle Area:", s.Area())
}
