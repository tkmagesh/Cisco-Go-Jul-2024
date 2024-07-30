package calculator

import "fmt"

// private (accessible across all the files within the package)
var opCount int

// package init function
func init() {
	fmt.Println("calculator[calc.go] - init")
}

// private
func demo() {
	fmt.Println("calculator demo")
}

// public
func OpCount() int {
	return opCount
}
