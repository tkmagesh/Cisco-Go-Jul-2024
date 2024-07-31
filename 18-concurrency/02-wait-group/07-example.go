package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10) // simulating a 'deadlock'
	go func() {
		defer wg.Done()
		f1()
	}()
	f2()
	wg.Wait()
}

// 3rd party api (to be executed as a goroutine)
func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
