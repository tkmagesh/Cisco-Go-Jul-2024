package main

import (
	"fmt"

	calc "github.com/tkmagesh/cisco-go-jul-2024/13-modularity-app/calculator"
	"github.com/tkmagesh/cisco-go-jul-2024/13-modularity-app/calculator/utils"
)

func main() {
	fmt.Println("modularity app executed")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Opcount :", calculator.OpCount())
	*/

	// using "calc" alias
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Opcount :", calc.OpCount())
	// calculator.demo()

	fmt.Println("IsEven(21) :", utils.IsEven(21))
	fmt.Println("IsOdd(21) :", utils.IsOdd(21))
}
