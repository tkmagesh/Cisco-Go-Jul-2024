package main

import (
	"fmt"
	"log"
)

func main() {

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	add := getLogOperation(add)
	subtract := getLogOperation(subtract)

	add(100, 200)
	subtract(100, 200)
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

func getLogOperation(operFn func(x, y int)) func(x, y int) {
	return func(x, y int) {
		log.Println("Invocation started")
		operFn(x, y)
		log.Println("Invocation completed")
	}
}
