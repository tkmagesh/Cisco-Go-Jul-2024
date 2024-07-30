package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Cost  float64
	Units int
}

type Products []Product

func (products Products) GetValue() float64 {
	var result float64
	for _, p := range products {
		result += p.Cost * float64(p.Units)
	}
	return result
}

func main() {
	products := Products{
		{101, "Pen", 10, 5},
		{102, "Pencil", 5, 20},
		{103, "Marker", 50, 2},
		{101, "Scribble Pad", 20, 10},
	}
	fmt.Println(products)
	fmt.Println(products.GetValue())
}
