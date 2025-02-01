package mypackage

import "fmt"

// Implicitly exported function by starting function name with a capital letter
func SayHello(name string) {
	fmt.Println("Hello,", name)
}
