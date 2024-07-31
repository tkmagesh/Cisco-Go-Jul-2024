package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for range 5 {
		fmt.Println(<-ch)
	}

}

func genNos(ch chan<- int) {
	for i := range 5 {
		ch <- (i + 1) * 10
		time.Sleep(500 * time.Millisecond)
	}

}
