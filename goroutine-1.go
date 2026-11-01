package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start printNumbers as a background task
	go printNumbers() 

	// The main function continues executing instantly
	fmt.Println("Started the goroutine!")
	
	// Wait long enough for the background task to finish before quitting
	time.Sleep(600 * time.Millisecond)
	fmt.Println("\nMain function finished.")
}
