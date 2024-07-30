package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("len(nos) =", len(nos))
	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// appending new values
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)

	// appending another slice
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// a slice is a pointer to an underlying array
	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	// slicing
	nos3 := nos[2:5] // from index(2) to index(5-1)
	fmt.Println(nos3)

	fmt.Println("Before sorting, nos =", nos)
	sort(nos)
	fmt.Println("After sorting, nos =", nos)

	// creating a NEW slice from another slice
	dupNos := make([]int, len(nos), len(nos)+1)
	copy(dupNos, nos)
	fmt.Println(dupNos)
}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
