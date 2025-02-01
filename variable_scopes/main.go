/*
	Variable scopes: Defines the part of the program where a variable is accessible.
*/

package main

import "fmt"

// Package-level (global) variable - accessible from anywhere in this package
var globalVar = "I'm a global variable"

func main() {
	// Function scope - accessible only within main()
	localVar := "I'm a local variable"

	// Block scope - accessible only within this if block
	if true {
		blockVar := "I'm a block-scoped variable"
		fmt.Println(blockVar) // Accessible here
	}

	fmt.Println(globalVar) // Accessible
	fmt.Println(localVar)  // Accessible

	// Uncommenting the next line will cause an error because blockVar is out of scope
	// fmt.Println(blockVar)

	// Call another function
	anotherFunction()
}

func anotherFunction() {
	// globalVar is accessible here
	fmt.Println(globalVar)

	// Uncommenting the next line will cause an error because localVar is not accessible here
	// fmt.Println(localVar)
}
