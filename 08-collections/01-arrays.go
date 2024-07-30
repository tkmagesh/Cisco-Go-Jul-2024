package main

import "fmt"

func main() {
	// var nos [5]int // memory is allocated & initialized with the zero value of the underlying data type
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos // a copy of nos is created coz arrays are values
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Printf("Before sorting, nos = %v\n", nos)
	sort(&nos)
	fmt.Printf("After sorting, nos = %v\n", nos)
}

func sort(list *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			/*
				if (*list)[i] > (*list)[j] {
					(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
				}
			*/
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
