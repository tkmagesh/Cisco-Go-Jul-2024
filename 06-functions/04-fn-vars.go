package main

import "fmt"

func main() {

	// 0 arguments, 0 return values
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	// 1 argument, 0 return values
	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	// 1 argument, 1 return value
	var greetUser func(string) string
	greetUser = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}
	msg := greetUser("Suresh")
	fmt.Println(msg)

	// 2 arguments, 1 return value
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	result := add(100, 200)
	fmt.Println("add result :", result)

	// 2 arguments, 2 return values
	// using named result(s)
	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) { // quotient & remainder are declared and initialized with the zero values of int
		quotient, remainder = x/y, x%y
		return
	}
	quotient, remainder := divide(100, 7)
	fmt.Printf("q = %d and r = %d\n", quotient, remainder)

}
