package mypackage

import "fmt"

// Implicitly exported function by starting function name with a capital letter
func SayBye(name string) {
	fmt.Println("Bye,", name)
}
