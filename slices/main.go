/*
	Slices: Similar to arrays but they are dynamic in size. They can grow or shrink.
*/

package main

import "fmt"

func main() {
	// Declare an empty slice with 0 length and 0 capacity
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1) // []

	// Declare a slice with 5 length and 5 capacity
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2) // Go Slices Are Powerful

	// Add elements to a slice
	myslice1 = append(myslice1, 1, 2, 3, 4, 5)
	fmt.Println(myslice1) // 1 2 3 4 5
}
