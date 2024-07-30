package calculator

import "fmt"

// private (accessible across all the files within the package)
// var opCount int

var opCount map[string]int

// package init function
func init() {
	opCount = make(map[string]int, 0)
	fmt.Println("calculator[calc.go] - init")
}

// private
func demo() {
	fmt.Println("calculator demo")
}

// public
/*
func OpCount() int {
	return opCount
}
*/

func OpCount() map[string]int {
	return opCount
}
