package calculator

import "fmt"

func init() {
	fmt.Println("calculator[subtract.go] - init")
}

func Subtract(x, y int) int {
	// opCount++
	opCount["Subtract"]++
	return x - y
}
