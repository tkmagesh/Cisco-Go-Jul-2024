package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count atomic.Int64

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		// modify to execute the increment() function concurrently
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count.Load())
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count.Add(1)

}
