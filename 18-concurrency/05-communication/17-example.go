package main

import "fmt"

// Receive Only Channel
// consumer
/*
func main() {

	ch := add(100, 200)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
*/

// Send Only Channel
// consumer
func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int, ch chan<- int) {
	result := x + y
	ch <- result
}
