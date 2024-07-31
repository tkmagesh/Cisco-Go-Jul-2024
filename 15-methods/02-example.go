/* methods inheritance & overriding */

package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

// a method == a function with a reciever
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

// override the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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
	milk := NewPerishableProduct(101, "Milk", 50, "2 Days")
	/*
		fmt.Println(milk.Format())
		milk.ApplyDiscount(10)
		fmt.Println(milk.Format())
	*/

	// fmt.Print functions will call the String() method if the object implements fmt.Stringer interface
	fmt.Println(milk)
	milk.ApplyDiscount(10)
	fmt.Println(milk)
}
