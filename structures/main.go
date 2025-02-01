/*
	Strucutres: Used to create a collection of members of different data
			    types into a single variable.
*/

package main

import "fmt"

// 1. Defining a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// 2. Creating a struct instance using a struct literal
	var p1 Person
	p1.Name = "Alice"
	p1.Age = 25
	fmt.Println("Person 1:", p1)

	// 3. Creating a struct instance using a short struct literal
	p2 := Person{"Bob", 30}
	fmt.Println("Person 2:", p2)

	// 4. Creating a struct instance using named fields
	p3 := Person{Name: "Charlie"}
	fmt.Println("Person 3:", p3)

	// 5. Creating a struct pointer
	p4 := &Person{"David", 40}
	fmt.Println("Person 4:", *p4) // Dereferencing to print actual struct

	// 6. Modifying struct fields via a pointer
	p4.Age = 41 // Dereferencing not required
	fmt.Println("Person 4 (updated age):", *p4)
}
