/*
	Interfaces: A type that lists methods without providing an implementation. An instance of an
				interface is created implicitly by implementing the methods of the interface.
*/

package main

import (
	"fmt"
	"math"
)

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type that implements the Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle type that implements the Shape interface
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Main function to demonstrate the interface
func main() {
	// Declare a variable of the interface type
	var s Shape

	// Assign the Circle type to the Shape interface
	s = Circle{radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	// Assign the Rectangle type to the Shape interface
	s = Rectangle{length: 4, width: 3}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
}
