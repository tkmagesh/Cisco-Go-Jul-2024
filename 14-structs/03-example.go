package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

// factory function for PerishableProduct (to hide the complexities of constructing the PerishableProduct instance)
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		milk := PerishableProduct{
			Product: Product{
				Id:   101,
				Name: "Milk",
				Cost: 50,
			},
			Expiry: "2 Days",
		}
	*/
	milk := NewPerishableProduct(101, "Milk", 50, "2 Days")
	fmt.Printf("%#v\n", milk)
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Product.Id, milk.Name, milk.Cost, milk.Expiry)
}
