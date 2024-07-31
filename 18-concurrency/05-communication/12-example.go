package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int) {
	defer wg.Done()
	result = x + y
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
