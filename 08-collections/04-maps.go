package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	/*
		var productRanks map[string]int = map[string]int{
			"pen":    3,
			"pencil": 1,
			"marker": 5,
		}
	*/

	/*
		var productRanks = map[string]int{
			"pen":    3,
			"pencil": 1,
			"marker": 5,
		}
	*/

	/*
		productRanks := map[string]int{
			"pen":    3,
			"pencil": 1,
			"marker": 5,
		}
	*/

	productRanks := make(map[string]int)
	productRanks["pen"] = 3
	productRanks["pencil"] = 1
	productRanks["marker"] = 5

	fmt.Println(productRanks)

	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iterating using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	// keyToCheck := "notepad"
	// productRanks["notepad"] = 0
	keyToCheck := "notepad"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("productRanks[%q] = %d\n", keyToCheck, val)
	} else {
		fmt.Printf("key [%v] does not exist!\n", keyToCheck)
	}

	// remove an key/value pair
	// keyToRemove := "pen"
	keyToRemove := "notepad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
