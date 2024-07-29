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

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// v2.0
/*
func logAdd(x, y int) {
	log.Println("Invocation started")
	add(x, y)
	log.Println("Invocation completed")
}

func logSubtract(x, y int) {
	log.Println("Invocation started")
	subtract(x, y)
	log.Println("Invocation completed")
}
*/

func logOperation(operFn func(x, y int), x, y int) {
	log.Println("Invocation started")
	operFn(x, y)
	log.Println("Invocation completed")
}
