package main

import (
	"fmt"
)

func main() {
	// Define two variables
	var a int = 5 // 0101 in binary
	var b int = 3 // 0011 in binary

	// Print variables in binary format
	fmt.Printf("a = %04b, b = %04b\n", a, b) // 0101, 0011

	// Bitwise AND
	fmt.Printf("a & b = %04b\n", a&b) // 0001

	// Bitwise OR
	fmt.Printf("a | b = %04b\n", a|b) // 0111

	// Bitwise XOR
	fmt.Printf("a ^ b = %04b\n", a^b) // 0110

	// Bitwise NOT
	fmt.Printf("^a = %04b\n", (^a)&15) // 1010

	// Left shift
	fmt.Printf("a << 1 = %04b\n", a<<1) // 1010

	// Right shift
	fmt.Printf("a >> 1 = %04b\n", a>>1) // 0010
}
