/*
	Maps: Used to store data in key-value pairs. A map is an unordered and
		  unchangeable collection that does not allow duplicates.
*/

package main

import "fmt"

func main() {
	// 1. Declare a map using make
	myMap := make(map[string]int)
	myMap["apple"] = 5
	fmt.Println("myMap:", myMap)

	// 2. Declare and initialize a map with values
	anotherMap := map[string]int{"banana": 2, "orange": 3}
	fmt.Println("anotherMap:", anotherMap)

	// 3. Access a value from a map
	value := myMap["apple"]
	fmt.Println("Value for 'apple':", value)

	// 4. Check if a key exists
	val, exists := myMap["grape"]
	if exists {
		fmt.Println("Value for 'grape':", val)
	} else {
		fmt.Println("'grape' not found in myMap")
	}

	// 5. Delete a key from the map
	delete(anotherMap, "banana")
	fmt.Println("anotherMap after deletion:", anotherMap)

	// 6. Iterate over a map
	fmt.Println("Iterating over myMap:")
	for key, val := range myMap {
		fmt.Println(key, "->", val)
	}
}
