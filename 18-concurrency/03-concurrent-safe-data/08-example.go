package main

import "fmt"

var count int

func main() {
	for range 300 {
		// modify to execute the increment() function concurrently
		increment()
	}
	fmt.Println("count :", count)
}

func increment() {
	count++
}
