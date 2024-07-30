package main

import "fmt"

func main() {
	fmt.Println("[main] started")
	defer func() {
		fmt.Println("	[main] deferred")
	}()
	f1()
	fmt.Println("[main] completed")
}

func f1() {
	fmt.Println("[f1] started")
	x := 100
	defer fmt.Println("	[f1] deferred - x :", x)
	/*
		defer func() {
			fmt.Println("	[f1] deferred - 1")
		}()
		defer func() {
			fmt.Println("	[f1] deferred - 2")
		}()
		defer func() {
			fmt.Println("	[f1] deferred - 3")
		}()
	*/

	defer fmt.Println("	[f1] deferred - 1")
	defer fmt.Println("	[f1] deferred - 2")
	defer fmt.Println("	[f1] deferred - 3")

	f2()
	fmt.Println("[f1] completed")
	x = 200
}

func f2() {
	fmt.Println("[f2] started")
	defer func() {
		fmt.Println("	[f2] deferred")
	}()
	fmt.Println("[f2] completed")
}
