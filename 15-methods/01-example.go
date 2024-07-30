package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {

	var pen *Product = &Product{
		Id:   100,
		Name: "Fountain Pen",
		Cost: 10,
	}

	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	fmt.Println("Applying Discount")
	/*
		fmt.Println("Before :", Format(pen))
		ApplyDiscount(&pen, 10)
		fmt.Println("Before :", Format(pen))
	*/
	fmt.Println("Before :", pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println("Before :", pen.Format())
}

// a method == a function with a reciever
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
