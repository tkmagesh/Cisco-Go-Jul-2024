package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for range 100 {
		fn := getFn()
		fn()
	}
}

func getFn() func() {
	if no := rand.Intn(100); no%2 == 0 {
		return f1
	} else {
		return f2
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
