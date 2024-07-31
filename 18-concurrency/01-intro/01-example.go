package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution of f1 through the scheduluer
	f2()

	// block the execution so that the scheduler will look for other goroutines that are scheduled and attempt to execute them
	// poor man's synchronization
	time.Sleep(4 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
