/*
Accept a range from the user and print all the prime numbers between the given range
*/

package main

import (
	"fmt"
	"sync"
)

var primes []int
var mutex sync.Mutex

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	// check each number if it is a prime CONCURRENTLY and print the prime numbers
	wg := &sync.WaitGroup{}
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(i) {
				mutex.Lock()
				{
					primes = append(primes, i)
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(primes)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
