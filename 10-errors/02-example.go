package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	muliplier := 100
	// divisor := 0
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divide(muliplier, divisor)
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

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

// using named results
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		// err = errors.New("divisor cannot be zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
