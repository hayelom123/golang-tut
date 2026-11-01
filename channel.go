package main

import "fmt"

func calculateSquare(num int, resultChan chan int) {
	resultChan <- num * num // Send the result into the channel
}

func main() {
	// Create a channel that transports integers
	myChannel := make(chan int)

	// Fire the goroutine
	go calculateSquare(5, myChannel)

	// Receive the value out of the channel (this blocks until data arrives)
	result := <-myChannel 

	fmt.Println("Result:", result)
}
