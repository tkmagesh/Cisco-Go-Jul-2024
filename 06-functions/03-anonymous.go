package main

import "fmt"

func main() {

	// 0 arguments, 0 return values
	func() {
		fmt.Println("Hi!")
	}()

	// 1 argument, 0 return values
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	// 1 argument, 1 return value
	msg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}("Suresh")
	fmt.Println(msg)

	// 2 arguments, 1 return value
	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("add result :", result)

	// 2 arguments, 2 return values
	// using named result(s)
	quotient, remainder := func(x, y int) (quotient, remainder int) { // quotient & remainder are declared and initialized with the zero values of int
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("q = %d and r = %d\n", quotient, remainder)

}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/
