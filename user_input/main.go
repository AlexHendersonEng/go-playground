package main

import "fmt"

func main() {
	// Declare a variable to store the user input
	var name string

	// Ask the user to enter their name and store in the 'name' variable
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // The '&' operator is used to get the memory address of the variable

	// Print user name
	fmt.Println("Hello", name)
}
