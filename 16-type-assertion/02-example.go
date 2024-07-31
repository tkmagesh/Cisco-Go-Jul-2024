package main

import (
	"fmt"
	"math"
)

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Neither a circle nor a rectangle!")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x does not have Area() method!")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

/*
To do:

	Introduce Perimeter() method to both Circle & Rectangle
	Write PrintPerimeter() function to print the perimeter of the given object
*/
func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{12, 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
	/*
		p := Product{100, "Pen", 10}
		PrintArea(p)
	*/
}
