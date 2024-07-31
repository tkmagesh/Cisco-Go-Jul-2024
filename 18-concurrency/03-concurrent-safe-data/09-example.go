package main

import (
	"fmt"
	"sync"
)

// encapsulate the count & mutex in a "concurrent safe" type

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter *Counter = &Counter{}

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		// modify to execute the increment() function concurrently
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
