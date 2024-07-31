package main

import (
	"fmt"
)

// consumer
func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}

/*
func main() {
	var result int
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200, &result)
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int, result *int) {
	defer wg.Done()
	*result = x + y
}
*/
