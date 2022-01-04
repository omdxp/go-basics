package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) perimeter() float64 {
	return 4 * s.length
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// print shape info
func printShapeInfo(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())
}

func main() {
	s := square{length: 5}
	c := circle{radius: 5}

	printShapeInfo(s)
	printShapeInfo(c)

	shapes := []shape{
		square{length: 5},
		circle{radius: 5},
		square{length: 10},
		circle{radius: 10},
	}

	for _, s := range shapes {
		printShapeInfo(s)
	}
}
