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

	no = 999
	fmt.Println("Before incrementing, no =", no)
	fmt.Println("&no =", &no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func increment(val *int) {
	fmt.Println("[increment] &val = ", val)
	*val++
}

func swap(x, y *int) /* no return values */ {
	*x, *y = *y, *x
}
