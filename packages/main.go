/*
	Packages: Allows for the grouping of related code.

	Modules: A collection of related packages that are versioned together
			 using a go.mod file. When building a package a go.sum file is
			 created to store the checksum of the dependencies.
*/

package main

import (
	"fmt"                 // Built-in package
	"mypackage/mypackage" // Custom package

	"github.com/fatih/color" // External package
)

func main() {
	// Using built-in fmt package
	fmt.Println("Using the built-in fmt package!")

	// Using our custom package
	mypackage.SayHello("Go Developer")
	mypackage.SayBye("Go Developer")

	// Using an external package (fatih/color)
	color.Green("This is a green message!")
}
