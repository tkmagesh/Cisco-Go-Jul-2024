package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	msg := getGreetMsg("Suresh")
	fmt.Println(msg)
	fmt.Println("add(100,200) =", add(100, 200))

	// fmt.Println(divide(100, 7))
	multiplier, divisor := 100, 7

	quotient, remainder := divide(multiplier, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, quotient, remainder)

	// Using _
	/*
		quotient, _ := divide(multiplier, divisor)
		fmt.Printf("Dividing %d by %d, quotient = %d \n", multiplier, divisor, quotient)
	*/
}

// 0 arguments, 0 return values
func sayHi() {
	fmt.Println("Hi!")
}

// 1 argument, 0 return values
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

// 1 argument, 1 return value
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a good day!", userName)
}

// 2 arguments, 1 return value
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

// 2 arguments, 2 return values
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using named result(s)
func divide(x, y int) (quotient, remainder int) { // quotient & remainder are declared and initialized with the zero values of int
	quotient, remainder = x/y, x%y
	return
}
