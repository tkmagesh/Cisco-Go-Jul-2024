/*
Accept a range from the user and print all the prime numbers between the given range
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	resultCh := primeDispatcher(start, end)
	for primeNo := range resultCh {
		fmt.Println("prime :", primeNo)
	}
	fmt.Println("Done")
}

func primeDispatcher(start, end int) <-chan int {
	resultCh := make(chan int)
	go func() {
		noCh := make(chan int)
		workerWg := &sync.WaitGroup{}
		for range 10 {
			workerWg.Add(1)
			go primeWorker(noCh, resultCh, workerWg)
		}
		for no := start; no <= end; no++ {
			noCh <- no
		}
		close(noCh)
		workerWg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func primeWorker(ch <-chan int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := range ch {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			resultCh <- no
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
