package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var pen Product
		// fmt.Println(pen)
		pen.Id = 100
		pen.Name = "Fountain Pen"
		pen.Cost = 10
		// fmt.Printf("%#v\n", pen)
		fmt.Printf("%+v\n", pen)
	*/

	var pen Product = Product{
		Id:   100,
		Name: "Fountain Pen",
		Cost: 10,
	}
	// fmt.Printf("%+v\n", pen)
	fmt.Println(Format(pen))

	var p2 Product = pen // a copy is created coz everything is a value
	p2.Cost = 20
	fmt.Printf("pen.Cost = %v, p2.Cost = %v\n", pen.Cost, p2.Cost)

	fmt.Println("Apply Discount")
	fmt.Println("Before :", pen)
	ApplyDiscount(&pen, 10)
	fmt.Println("Before :", pen)
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
