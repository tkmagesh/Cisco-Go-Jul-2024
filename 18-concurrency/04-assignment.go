/*
Accept a range from the user and print all the prime numbers between the given range
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	// check each number if it is a prime CONCURRENTLY and print the prime numbers
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
