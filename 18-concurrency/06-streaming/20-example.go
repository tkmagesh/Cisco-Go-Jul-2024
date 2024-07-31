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
	for {
		time.Sleep(500 * time.Millisecond)
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		fmt.Println("Consumer done!")
		break
	}

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
