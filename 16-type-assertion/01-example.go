package main

import "fmt"

func main() {
	var x interface{}
	/*
		x = 100
		x = true
		x = 99.99
		x = struct{}{}
		fmt.Println(x)
	*/
	// x = 100
	x = "Ut ea dolor excepteur cupidatat sint nulla."
	// y := x.(int) + 200

	// type assertion - 1
	/*
		// x = "Ut ea dolor excepteur cupidatat sint nulla."
		x = 100
		if val, ok := x.(int); ok {
			y := val + 200
			fmt.Println(y)
		} else {
			fmt.Println("x is not an int")
		}
	*/

	// type assertion - 2
	// x = 100
	// x = "Cillum irure tempor duis sit aliquip."
	// x = true
	// x = 99.99
	// x = struct{}{}
	x = struct{ id int }{id: 100}
	switch val := x.(type) {
	case int:
		fmt.Printf("x is an int, x + 200 = %d\n", val+200)
	case string:
		fmt.Printf("x is a string, len(x) = %d\n", len(val))
	case bool:
		fmt.Printf("x is a bool, !x = %v\n", !val)
	case float64:
		fmt.Printf("x is a float64, 0.01%% of x = %v\n", val*0.01)
	case struct{}:
		fmt.Println("x is an 0 byte struct")
	default:
		fmt.Println("x is an unknown type!")
	}
}
