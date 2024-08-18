package main

import (
	"fmt"
	"math"
)

// Shape is an interface with methods for calculating area and perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is a struct representing a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a struct representing a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// CalculateShapeInfo takes a Shape interface and calculates its area and perimeter.
func CalculateShapeInfo(s Shape) {
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Printf("Shape Info:\nArea: %.2f\nPerimeter: %.2f\n", area, perimeter)
}

func main() {
	// Create a rectangle
	rect := Rectangle{Width: 5, Height: 3}

	// Create a circle
	circ := Circle{Radius: 2}

	// Calculate and display information for the rectangle
	fmt.Println("Rectangle:")
	CalculateShapeInfo(rect)
	fmt.Println()

	// Calculate and display information for the circle
	fmt.Println("Circle:")
	CalculateShapeInfo(circ)
}
