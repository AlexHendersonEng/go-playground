/*
	Data Types: The basic data types in Go are:
		- bool
		- string
		- int, int8, int16, int32, int64
		- uint, uint8, uint16, uint32, uint64, uintptr
		- byte (alias for uint8)
		- rune (alias for int32)
		- float32, float64
		- complex64, complex128
*/

package main

import "fmt"

func main() {
	// Boolean
	var isTrue bool = true
	fmt.Println(isTrue)

	// String
	var message string = "Hello, World!"
	fmt.Println(message)

	// Integers
	var number int = 42
	fmt.Println(number)

	// Floats
	var pi float32 = 3.14159
	fmt.Println(pi)

	// Complex
	var complexNumber complex64 = 1 + 2i
	fmt.Println(complexNumber)
}
