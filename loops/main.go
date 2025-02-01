/*
	Loops: Loop through a block of code a specified number of times.
*/

package main

import "fmt"

func main() {
	// 1. Basic for loop (like a C-style loop)
	for i := 0; i < 5; i++ {
		fmt.Println("Basic for loop:", i)
	}

	// 2. For loop as a while loop (condition only)
	j := 0
	for j < 5 {
		fmt.Println("While-style loop:", j)
		j++
	}

	// 3. Infinite loop with break
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("Infinite loop with break:", k)
		k++
	}

	// 4. Range-based loop (iterating over an array)
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Range-based loop:", index, value)
	}
}
