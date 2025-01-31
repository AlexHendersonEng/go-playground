/*
	Variables: Containers for storing data values.
*/

package main

import "fmt"

func main() {
	// Declaring a variable using the 'var' keyword. Requires the specification
	// of either the type or value or both. If no value is defined the variable will
	// be initialised with the default value of the type. Can be used inside and outside
	// of functions.
	var name string = "John Doe"
	fmt.Println(name) // John Doe

	// Declaring a variable with the ':=' operator. The type is inferred from the value.
	// Can only be used inside functions.
	age := 25
	fmt.Println(age) // 25

	// Declaring multiple variables in one line
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a, b, c, d) // 1 3 5 7
	var e, f = 9, "Hello"
	fmt.Println(e, f) // 9 Hello
	g, h := 11, "World!"
	fmt.Println(g, h) // 11 World!

	// Declaring multiple variables in a block
	var (
		i int    = 1
		j int    = 2
		k string = "Bye"
	)
	fmt.Println(i, j, k) // 1 2 Bye
}
