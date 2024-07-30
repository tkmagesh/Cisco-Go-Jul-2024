package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred, err :", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	muliplier := 100
	divisor := 0
	q, r := divide(muliplier, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", muliplier, divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) { // quotient & remainder are declared and initialized with the zero values of int
	fmt.Println("Calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
