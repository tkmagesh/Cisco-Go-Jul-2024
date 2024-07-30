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
	// divisor := 0
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideClient(muliplier, divisor)
		if err == nil {
			fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", muliplier, divisor, q, r)
			break
		} else if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by zero")
		} else {
			fmt.Println("Unknown Error :", err)
		}
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		// conver the panic into an error
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (raises a panic)
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
