package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the wg counter by 1
	go f1()   // scheduling the execution of f1 through the scheduluer
	f2()

	// block the execution so that the scheduler will look for other goroutines that are scheduled and attempt to execute them
	wg.Wait() // block until the waitgroup counter becomes 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
