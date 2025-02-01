/*
	Goroutines: Allows functions to run concurrently using less memory than traditional threads.
	            Every Go program starts with a main Goroutine and if it exits all others stop.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Define functions
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	// 1. Basic Goroutine
	go sayHello()

	// 2. Anonymous function Goroutine
	go func() {
		fmt.Println("Hello from an anonymous goroutine!")
	}()

	// 3. Goroutine with Arguments
	go sayMessage("Hello with an argument!")

	// 4. Using WaitGroup
	var wg sync.WaitGroup
	wg.Add(2) // Increases the WaitGroup counter by 2

	go func() {
		defer wg.Done() // Decreases the WaitGroup counter by 1
		fmt.Println("First goroutine with WaitGroup!")
	}()

	go func() {
		defer wg.Done() // Decreases the WaitGroup counter by 1
		fmt.Println("Second goroutine with WaitGroup!")
	}()

	// 5. Using Channels
	msgChan := make(chan string) // Create a channel of type string

	go func() {
		msgChan <- "Hello from a channel!" // Sending message to the channel using send operator
	}()

	fmt.Println(<-msgChan) // Receiving message from the channel using receive operator

	wg.Wait() // Wait for goroutines until WaitGroup counter is 0

	time.Sleep(100 * time.Millisecond) // Give some time for non-WaitGroup goroutines to finish
}
