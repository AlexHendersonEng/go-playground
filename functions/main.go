/*
	Functions: A block of statments that can be used repeatedly in a program by calling the function.
*/

package main

import "fmt"

// 1. Basic function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello!")
}

// 2. Function with parameters
func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// 3. Function with parameters and a return value
func multiply(a int, b int) int {
	return a * b
}

// 4. Function with multiple return values
func divide(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// 5. Function with named return values
func rectangleArea(length int, width int) (area int) {
	area = length * width // No need for explicit return statement
	return
}

// 6. Anonymous function (immediately invoked)
func main() {
	// Call function with no parameters
	sayHello()

	// Call function with parameters
	add(3, 5)

	// Call function with parameters and return value
	result := multiply(4, 6)
	fmt.Println("Product:", result)

	// Call function with multiple return values
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// Call function with named return values
	rectArea := rectangleArea(5, 10)
	fmt.Println("Rectangle Area:", rectArea)

	// Anonymous function assigned to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Alice")

	// Immediately invoked anonymous function
	func() {
		fmt.Println("This runs immediately!")
	}()
}
