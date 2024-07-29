package main

import "fmt"

func main() {
	var no int = 100

	var noPtr *int
	noPtr = &no // value => address
	fmt.Println(noPtr)

	// address => value (dereferencing)
	var x = *noPtr
	fmt.Println(x)

	no = 200
	x = *noPtr
	fmt.Println(x)
}
