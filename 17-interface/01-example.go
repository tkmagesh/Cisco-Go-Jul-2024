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

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

/*
To do:

	Introduce Perimeter() method to both Circle & Rectangle
	Write PrintPerimeter() function to print the perimeter of the given object
*/
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)      // x => interface{ Area() float64 }
	PrintPerimeter(x) // x => interface { Perimeter() float64}
}

/*
	func PrintStats(x interface {
		Area() float64
		Perimeter() float64
	}) {

		PrintArea(x)      // x => interface{ Area() float64 }
		PrintPerimeter(x) // x => interface { Perimeter() float64}
	}
*/
func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{12, 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

	/*
		p := Product{100, "Pen", 10}
		PrintArea(p)
	*/
}
