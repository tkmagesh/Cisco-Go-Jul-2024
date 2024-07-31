package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int, ch chan int) {
	defer wg.Done()
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
