/*
	Type keyword: The 'type' keyword is used to define a new type based on an existing type.
*/

package main

import "fmt"

// Define a new type called 'Age' based on int
type Age int

func main() {
	// Create a variable of type 'Age'
	var myAge Age = 25

	// Print the value of 'myAge'
	fmt.Println("My age is:", myAge) // My age is: 25
}
