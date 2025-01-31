/*
	Arrays: Used to stor multiple values of the same type in a single variable.
*/

package main

import "fmt"

func main() {
	// Declare an array with the 'var' keyword
	var array1 = [5]int{1, 2, 3, 4, 5}   // Explicitly declare the size of the array
	fmt.Println(array1)                  // 1 2 3 4 5
	var array2 = [...]int{1, 2, 3, 4, 5} // Implicitly declare the size of the array
	fmt.Println(array2)                  // 1 2 3 4 5

	// Declare an array using the ':=' operator
	array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3) // 1 2 3 4 5

	// Array initialisation
	array4 := [5]int{}              // Not initialised
	fmt.Println(array4)             // 0 0 0 0 0
	array5 := [5]int{1, 2}          // Partially initialised
	fmt.Println(array5)             // 1 2 0 0 0
	array6 := [5]int{1, 2, 3, 4, 5} // Fully initialised
	fmt.Println(array6)             // 1 2 3 4 5
	array7 := [5]int{1: 10, 3: 30}  // Partially initialised (element 2 initialised to 10 and element 4 initialised to 30)
	fmt.Println(array7)             // 0 10 0 30 0

	// Change element in an array using an index
	array1[2] = 10
	fmt.Println(array1) // 1 2 10 4 5
}
