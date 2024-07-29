package main

import "fmt"

// var app_name string = "Variables_App"

// type inference
var app_name = "Variables_App"

// the following will not work
// app_name := "Variables_App"

func main() {
	/*
		var userName string
		userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	/*
		var userName string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// type inference
	/*
		var userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// using :=
	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	// using package scoped variable
	// fmt.Printf("App Name : %q\n", app_name)

	var app_version = "2.0"
	app_version = "3.0"
	fmt.Printf("app version : %s\n", app_version) // reading app_version and there by making it 'in use'

	// Multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x,y int = 100, 200
		var str string = "sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// using ":="

	x, y, str := 100, 200, "sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
