/*
	switch statements: A switch statement is a control structure that allows you to test
					   a variable against a list of values. Each value is called a case,
					   and the variable being switched on is checked for each case.
*/

package main

import (
	"fmt"
)

func main() {
	// Basic switch
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}

	// Switch with multiple cases
	letter := "b"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	// Switch without an expression (like an if-else chain)
	age := 18
	switch {
	case age < 18:
		fmt.Println("Minor")
	case age == 18:
		fmt.Println("Just turned adult")
	default:
		fmt.Println("Adult")
	}

	// Type switch (for interface{} types)
	var x interface{} = 10.5 // interface{} is a type that can hold values of any type
	switch v := x.(type) {
	case int:
		fmt.Println("x is an int:", v)
	case float64:
		fmt.Println("x is a float64:", v)
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
