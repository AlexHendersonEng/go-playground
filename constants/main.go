/*
	Constants: Variables declared using the 'const' keyword are called constants. This means
			   they are unchangeable and read only.
*/

package main

import "fmt"

func main() {
	// Define a constant with explicit and implicit type
	const CONST1 string = "This is a constant"
	const CONST2 = 100
	fmt.Println(CONST1, CONST2) // This is a constant 100
}
