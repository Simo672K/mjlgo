package main

import "fmt"

func main() {
	// for loops in golang
	for i := 0; i < 5; i++ {
		fmt.Println("iteration number: ", i)
	}

	// to work with slices
	mySlice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("at position (%d): %d\n", i, mySlice[i])
	}

	// another way to work with slices
	for i := range mySlice {
		fmt.Println(i) // i here is only the index, to show the value ethier mySlice[i] or distructuring it in loop declaration
	}

	// getting value
	for i, value := range mySlice {
		fmt.Println(i, value)
		if i == 1 {
			break // break of the loop
		}
	}

	// getting value
	for _, value := range mySlice { // _ to ignore the index if we don't need it
		fmt.Println(value)
	}
}
