package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Consumer done!")
}

// producer
func genNos(ch chan<- int) {
	var i int
	for ; ; i++ {
		ch <- (i + 1) * 10
		if rand.Intn(100)%3 == 0 {
			fmt.Println("All the numbers are produced")
			close(ch)
			break
		}
	}
	fmt.Println("Producer done!")
}
